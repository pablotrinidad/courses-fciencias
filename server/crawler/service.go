// package crawler exposes the gRPC implementation of service FCCrawler.
package crawler

import (
	"context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// FCCrawlerServiceImpl is an implementation of FCCrawler gRPC service.
type FCCrawlerServiceImpl struct{}

// NewFCCrawlerServiceImpl returns a pointer to the service implementation.
func NewFCCrawlerServiceImpl() *FCCrawlerServiceImpl {
	return &FCCrawlerServiceImpl{}
}

// ListMajors return all available majors. It will fetch majors from website even though data is
// known and unlikely to change.
func (*FCCrawlerServiceImpl) ListMajors(ctx context.Context, req *ListMajorsRequest) (*ListMajorsResponse, error) {
	type fetchMajorResult struct {
		m   *major
		err error
	}

	cn := make(chan *fetchMajorResult, len(majorsList))
	fns := make([]func(), 0, len(majorsList))
	for majorID, _ := range majorsList {
		fn := func() {
			major, err := fetchMajor(majorID)
			cn <- &fetchMajorResult{m: major, err: err}
		}
		fns = append(fns, fn)
	}
	callConcurrent(fns...)
	close(cn)

	resp := &ListMajorsResponse{}
	resp.Majors = make([]*Major, 0, len(majorsList))

	for r := range cn {
		if r.err != nil {
			return nil, status.Errorf(codes.Internal, "an error occurred processing major %d", r.m.externalID)
		}
		resp.Majors = append(resp.Majors, r.m.toProto())
	}

	return resp, nil
}
