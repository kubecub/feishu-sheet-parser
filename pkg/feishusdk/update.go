// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package feishusdk

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

// UpdateBitableMeta 通过 app_token 更新多维表格元数据
// 飞书文档、飞书表格、知识库中的多维表格不支持开启高级权限
// 此接口非原子操作, 先修改多维表格名字, 后开关高级权限。可能存在部分成功的情况
// 该接口支持调用频率上限为 10 QPS
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/update
func (cli *BitableClient) UpdateMeta(ctx context.Context, baseToken string, name *string, isAdvanced *bool) (*lark.UpdateBitableMetaRespApp, error) {
	res, response, err := cli.larkCli.Bitable.UpdateBitableMeta(ctx, &lark.UpdateBitableMetaReq{
		Name:       name,
		AppToken:   baseToken,
		IsAdvanced: isAdvanced,
	})
	if err != nil {
		return nil, fmt.Errorf("UpdateBitableMeta failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.App, nil
}

// UpdateBitableField 该接口用于在数据表中更新一个字段
// 该接口支持调用频率上限为 10 QPS
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/update
func (cli *BitableClient) UpdateField(ctx context.Context, request *lark.UpdateBitableFieldReq) (*lark.UpdateBitableFieldRespField, error) {
	res, response, err := cli.larkCli.Bitable.UpdateBitableField(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("UpdateBitableField failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.Field, nil
}

// UpdateBitableRecord 该接口用于更新数据表中的一条记录
// 该接口支持调用频率上限为 10 QPS
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/update
func (cli *BitableClient) UpdateRecord(ctx context.Context, baseToken string, tableID string, recordID string, fields map[string]interface{}) (*lark.UpdateBitableRecordRespRecord, error) {
	userIDType := lark.IDTypeUserID
	res, response, err := cli.larkCli.Bitable.UpdateBitableRecord(ctx, &lark.UpdateBitableRecordReq{
		UserIDType: &userIDType,
		AppToken:   baseToken,
		RecordID:   recordID,
		TableID:    tableID,
		Fields:     fields,
	})
	if err != nil {
		return nil, fmt.Errorf("UpdateBitableRecord failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.Record, nil
}

// BatchUpdateBitableRecord 该接口用于更新数据表中的多条记录, 单词调用最多更新 500 条记录。
// 该接口支持调用频率上限为 10 QPS
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_update
func (cli *BitableClient) BatchUpdateRecord(ctx context.Context, baseToken string, tableID string, recordList []*lark.BatchUpdateBitableRecordReqRecord) ([]*lark.BatchUpdateBitableRecordRespRecord, error) {
	userIDType := lark.IDTypeUserID
	res, response, err := cli.larkCli.Bitable.BatchUpdateBitableRecord(ctx, &lark.BatchUpdateBitableRecordReq{
		UserIDType: &userIDType,
		AppToken:   baseToken,
		Records:    recordList,
		TableID:    tableID,
	})
	if err != nil {
		return nil, fmt.Errorf("BatchUpdateBitableRecord failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.Records, nil
}
