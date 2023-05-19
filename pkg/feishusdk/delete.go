package feishusdk

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

//DeleteBitableTable 删除一个数据表
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/delete
func (cli *BitableClient) DeleteTable(ctx context.Context, baseToken string, tableID string) error {
	_, response, err := cli.larkCli.Bitable.DeleteBitableTable(ctx, &lark.DeleteBitableTableReq{
		AppToken: baseToken,
		TableID:  tableID,
	})
	if err != nil {
		return fmt.Errorf("DeleteBitableTable failed,  response:%+#v, err=%w", response, err)
	}
	return nil
}

//BatchDeleteBitableTable 删除多个数据表
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/batch_delete
func (cli *BitableClient) BatchDeleteTable(ctx context.Context, baseToken string, tableIDList []string) error {
	_, response, err := cli.larkCli.Bitable.BatchDeleteBitableTable(ctx, &lark.BatchDeleteBitableTableReq{
		AppToken: baseToken,
		TableIDs: tableIDList,
	})
	if err != nil {
		return fmt.Errorf("BatchDeleteBitableTable failed,  response:%+#v, err=%w", response, err)
	}
	return nil
}

//DeleteBitableField 该接口用于在数据表中删除一个字段
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/delete
func (cli *BitableClient) DeleteField(ctx context.Context, baseToken string, tableID string, fieldID string) error {
	_, response, err := cli.larkCli.Bitable.DeleteBitableField(ctx, &lark.DeleteBitableFieldReq{
		AppToken: baseToken,
		TableID:  tableID,
		FieldID:  fieldID,
	})
	if err != nil {
		return fmt.Errorf("DeleteBitableField failed,  response:%+#v, err=%w", response, err)
	}
	return nil
}

//DeleteBitableView 删除数据表中的视图
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/delete
func (cli *BitableClient) DeleteView(ctx context.Context, baseToken string, tableID string, viewID string) error {
	_, response, err := cli.larkCli.Bitable.DeleteBitableView(ctx, &lark.DeleteBitableViewReq{
		AppToken: baseToken,
		TableID:  tableID,
		ViewID:   viewID,
	})
	if err != nil {
		return fmt.Errorf("DeleteBitableView failed,  response:%+#v, err=%w", response, err)
	}
	return nil
}

//DeleteBitableRecord 该接口用于删除数据表中的一条记录
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/delete
func (cli *BitableClient) DeleteRecord(ctx context.Context, baseToken string, tableID string, recordID string) error {
	_, response, err := cli.larkCli.Bitable.DeleteBitableRecord(ctx, &lark.DeleteBitableRecordReq{
		AppToken: baseToken,
		TableID:  tableID,
		RecordID: recordID,
	})
	if err != nil {
		return fmt.Errorf("DeleteBitableRecord failed,  response:%+#v, err=%w", response, err)
	}
	return nil
}

//BatchDeleteBitableRecord 该接口用于删除数据表中现有的多条记录, 单次调用中最多删除 500 条记录。
//该接口支持调用频率上限为 10 QPS
//doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_delete
func (cli *BitableClient) BatchDeleteRecord(ctx context.Context, baseToken string, tableID string, recordIDList []string) error {
	_, response, err := cli.larkCli.Bitable.BatchDeleteBitableRecord(ctx, &lark.BatchDeleteBitableRecordReq{
		AppToken: baseToken,
		TableID:  tableID,
		Records:  recordIDList,
	})
	if err != nil {
		return fmt.Errorf("BatchDeleteBitableRecord failed,  response:%+#v, err=%w", response, err)
	}
	return nil
}