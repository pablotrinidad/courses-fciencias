// package crawler exposes the gRPC implementation of service FCCrawler.
package crawler

import (
	"context"
	"fmt"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// FCCrawlerServiceImpl is an implementation of FCCrawler gRPC service.
type FCCrawlerServiceImpl struct{}

// NewFCCrawlerServiceImpl returns a pointer to the service implementation.
func NewFCCrawlerServiceImpl() *FCCrawlerServiceImpl {
	return &FCCrawlerServiceImpl{}
}

// ListMajors returns all available majors. It will fetch majors from website even though data is
// known and unlikely to change.
func (*FCCrawlerServiceImpl) ListMajors(context.Context, *ListMajorsRequest) (*ListMajorsResponse, error) {
	majors, err := fetchMajorsConcurrently()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	resp := &ListMajorsResponse{Count: uint32(len(majors))}
	resp.Majors = make([]*Major, len(majors))
	for i, m := range majors {
		resp.Majors[i] = m.toProto()
	}
	return resp, nil
}

// TODO: Make method fetch available programs (must filter out those overlapping and without info)
// ListPrograms returns all available programs grouped by major. Data returned by this function is
// hardcoded since there is nothing more to fetch.
func (*FCCrawlerServiceImpl) ListPrograms(context.Context, *ListProgramsRequest) (*ListProgramsResponse, error) {
	majors, err := fetchMajorsConcurrently()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)

	}

	programs, err := fetchProgramsConcurrently()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	resp := &ListProgramsResponse{}
	resp.Majors = make([]*ListProgramsResponse_MajorBreakdown, len(majors))
	var programsCount uint32
	for i, _ := range majors {
		m := majors[i]
		mb := &ListProgramsResponse_MajorBreakdown{Major: m.toProto()}
		mPrograms := programs[m.externalID]
		for j, _ := range mPrograms {
			p := mPrograms[j]
			mb.Programs = append(mb.Programs, p.toProto())
			programsCount++
		}
		resp.Majors[i] = mb
	}
	resp.Count = programsCount
	return resp, nil
}

// fetchMajorsConcurrently visit majors' website and retrieves name.
func fetchMajorsConcurrently() ([]*major, error) {
	type fetchMajorResult struct {
		m   *major
		err error
	}

	cn := make(chan *fetchMajorResult, len(majorsList))
	fns := make([]func(), len(majorsList))
	for i, _ := range majorsList {
		majorID := majorsList[i]
		fns[i] = func() {
			major, err := fetchMajor(majorID)
			cn <- &fetchMajorResult{m: major, err: err}
		}
	}
	callConcurrent(fns)
	close(cn)

	majors := make([]*major, 0, len(majorsList))
	for r := range cn {
		if r.err != nil {
			return nil, fmt.Errorf("an error occurred processing major %d", r.m.externalID)
		}
		majors = append(majors, r.m)
	}
	return majors, nil
}

// fetchProgramsConcurrently returns a map indexed by major ID with the programs data.
func fetchProgramsConcurrently() (map[int][]*program, error) {
	type fetchProgramResult struct {
		majorID int
		p       *program
		err     error
	}
	cn := make(chan *fetchProgramResult)
	fns := make([]func(), 0)
	for mID, ps := range programs {
		for i := range ps {
			pID := ps[i]
			m := mID
			f := func() {
				p, err := fetchProgram(m, pID)
				cn <- &fetchProgramResult{majorID: m, p: p, err: err}
			}
			fns = append(fns, f)
		}
	}
	go func() {
		callConcurrent(fns)
		close(cn)
	}()

	programs := make(map[int][]*program)
	for r := range cn {
		if r.err != nil {
			return nil, r.err
		}
		programs[r.majorID] = append(programs[r.majorID], r.p)
	}
	return programs, nil
}
