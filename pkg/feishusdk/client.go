// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package feishusdk

import (
	"errors"

	"github.com/chyroc/lark"
)

type BitableClient struct {
	larkCli *lark.Lark
}

func NewClient(appID string, appSecret string) (*BitableClient, error) {
	if appID == "" || appSecret == "" {
		return nil, errors.New("invalid app credentials")
	}
	return &BitableClient{
		larkCli: lark.New(lark.WithAppCredential(appID, appSecret)),
	}, nil
}

func MustNewClient(appID string, appSecret string) *BitableClient {
	cli, err := NewClient(appID, appSecret)
	if err != nil {
		panic(err)
	}
	return cli
}
