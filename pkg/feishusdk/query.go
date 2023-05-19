package feishusdk

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

// GetBitableMeta 通过 app_token 获取多维表格元数据
// 该接口支持调用频率上限为 20 QPS
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/get
func (cli *BitableClient) GetMeta(ctx context.Context, baseToken string) (*lark.GetBitableMetaRespApp, error) {
	res, response, err := cli.larkCli.Bitable.GetBitableMeta(ctx, &lark.GetBitableMetaReq{
		AppToken: baseToken,
	})
	if err != nil {
		return nil, fmt.Errorf("GetBitableMeta failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.App, nil
}

// GetBitableTableList 根据 app_token, 获取多维表格下的所有数据表
// 该接口支持调用频率上限为 20 QPS
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/list
func (cli *BitableClient) GetTableList(ctx context.Context, baseToken string, pageSize *int64, pageToken *string) (*lark.GetBitableTableListResp, error) {
	res, response, err := cli.larkCli.Bitable.GetBitableTableList(ctx, &lark.GetBitableTableListReq{
		AppToken:  baseToken,
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		return nil, fmt.Errorf("GetBitableViewList failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res, nil
}

// GetBitableViewList 根据 app_token 和 table_id, 获取数据表的所有视图
// 该接口支持调用频率上限为 20 QPS
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/list
func (cli *BitableClient) GetViewList(ctx context.Context, baseToken string, tableID string, pageSize *int64, pageToken *string) (*lark.GetBitableViewListResp, error) {
	res, response, err := cli.larkCli.Bitable.GetBitableViewList(ctx, &lark.GetBitableViewListReq{
		AppToken:  baseToken,
		TableID:   tableID,
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		return nil, fmt.Errorf("GetBitableViewList failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res, nil
}

// GetBitableRecordList 该接口用于列出数据表中的现有记录, 单次最多列出 500 行记录, 支持分页获取。
// 该接口支持调用频率上限为 1000 次/分钟
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/list
func (cli *BitableClient) GetRecordList(ctx context.Context, request *lark.GetBitableRecordListReq) (*lark.GetBitableRecordListResp, error) {
	res, response, err := cli.larkCli.Bitable.GetBitableRecordList(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("GetBitableRecordList failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res, nil
}

// GetBitableRecord 该接口用于根据 record_id 的值检索现有记录
// 该接口支持调用频率上限为 20 QPS
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/get
func (cli *BitableClient) GetRecord(ctx context.Context, request *lark.GetBitableRecordReq) (*lark.GetBitableRecordResp, error) {
	res, response, err := cli.larkCli.Bitable.GetBitableRecord(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("GetBitableRecord failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res, nil
}