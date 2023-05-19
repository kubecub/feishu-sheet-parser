package feishusdk

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

//CreateBitableTable 新增一个数据表
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/create
func (cli *BitableClient) CreateTable(ctx context.Context, baseToken string, tableReq *lark.CreateBitableTableReqTable) (tableID string, err error) {
	userIDType := lark.IDTypeUserID
	res, response, err := cli.larkCli.Bitable.CreateBitableTable(ctx, &lark.CreateBitableTableReq{
		AppToken:   baseToken,
		Table:      tableReq,
		UserIDType: &userIDType,
	})
	if err != nil {
		return "", fmt.Errorf("CreateBitableTable failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.TableID, nil
}

//BatchCreateBitableTable 新增多个数据表
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/batch_create
func (cli *BitableClient) BatchCreateTable(ctx context.Context, baseToken string, tableReq []*lark.BatchCreateBitableTableReqTable) (tableIDs []string, err error) {
	userIDType := lark.IDTypeUserID
	res, response, err := cli.larkCli.Bitable.BatchCreateBitableTable(ctx, &lark.BatchCreateBitableTableReq{
		AppToken:   baseToken,
		Tables:     tableReq,
		UserIDType: &userIDType,
	})
	if err != nil {
		return nil, fmt.Errorf("BatchCreateBitableTable failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.TableIDs, nil
}

//CreateBitableRecord 该接口用于在数据表中新增一条记录
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/create
func (cli *BitableClient) CreateRecord(ctx context.Context, baseToken string, tableID string, fields map[string]interface{}) (*lark.CreateBitableRecordRespRecord, error) {
	userIDType := lark.IDTypeUserID
	res, response, err := cli.larkCli.Bitable.CreateBitableRecord(ctx, &lark.CreateBitableRecordReq{
		UserIDType: &userIDType,
		AppToken:   baseToken,
		TableID:    tableID,
		Fields:     fields,
	})
	if err != nil {
		return nil, fmt.Errorf("CreateBitableRecord failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.Record, nil
}

//BatchCreateBitableRecord 该接口用于在数据表中新增多条记录, 单次调用最多新增 500 条记录
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_create
func (cli *BitableClient) BatchCreateRecord(ctx context.Context, baseToken string, tableID string, recordsList []*lark.BatchCreateBitableRecordReqRecord) ([]*lark.BatchCreateBitableRecordRespRecord, error) {
	userIDType := lark.IDTypeUserID
	res, response, err := cli.larkCli.Bitable.BatchCreateBitableRecord(ctx, &lark.BatchCreateBitableRecordReq{
		UserIDType: &userIDType,
		AppToken:   baseToken,
		TableID:    tableID,
		Records:    recordsList,
	})
	if err != nil {
		return nil, fmt.Errorf("BatchCreateBitableRecord failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.Records, nil
}

//CreateBitableView 在数据表中新增一个视图
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/create
func (cli *BitableClient) CreateView(ctx context.Context, baseToken string, tableID string, viewName string, viewType *string) (*lark.CreateBitableViewRespView, error) {
	res, response, err := cli.larkCli.Bitable.CreateBitableView(ctx, &lark.CreateBitableViewReq{
		AppToken: baseToken,
		TableID:  tableID,
		ViewName: viewName,
		ViewType: viewType,
	})
	if err != nil {
		return nil, fmt.Errorf("CreateBitableView failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.View, nil
}

//CreateBitableField 该接口用于在数据表中新增一个字段
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/create
func (cli *BitableClient) CreateField(ctx context.Context, request *lark.CreateBitableFieldReq) (*lark.CreateBitableFieldRespField, error) {
	res, response, err := cli.larkCli.Bitable.CreateBitableField(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("CreateBitableField failed, res:%+v, response:%+#v, err=%w", res, response, err)
	}
	return res.Field, nil
}