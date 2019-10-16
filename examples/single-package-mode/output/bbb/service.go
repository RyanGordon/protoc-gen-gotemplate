// file generated with protoc-gen-gotemplate
package bbb

import (
	"fmt"

	"golang.org/x/net/context"
	"github.com/RyanGordon/protoc-gen-gotemplate/examples/single-package-mode/output/aaa"
)

type Service struct{}

func (service Service) Aaa(ctx context.Context, input *aaa.AaaRequest) (*aaa.AaaReply, error) {
	return nil, fmt.Errorf("method not implemented")
}

func (service Service) Bbb(ctx context.Context, input *BbbRequest) (*BbbReply, error) {
	return nil, fmt.Errorf("method not implemented")
}
