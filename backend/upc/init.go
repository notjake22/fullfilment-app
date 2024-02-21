package upc

import (
	"context"
)

func InitializeUPCLookup(upc string, ctx context.Context) (*LookUp, error) {
	u := &LookUp{
		UPC:     upc,
		item:    ItemData{},
		Context: ctx,
	}

	//err := u.proxyHandler()
	//if err != nil {
	//	return nil, err
	//}
	//
	return u, nil
}
