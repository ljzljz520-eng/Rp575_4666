package report

var statusLabels = map[string]string{
	"received": "已收货", "inspecting": "质检中", "accepted": "已验收", "rejected": "已拒收", "open": "待结算", "approved": "已批准", "paid": "已支付", "void": "已作废",
	"draft": "草稿", "active": "启用", "inactive": "停用", "pending": "待处理", "exported": "已导出", "logged": "已记录", "blocked": "已阻断", "expired": "已过期",
	"warehouse-a": "一号仓", "warehouse-b": "二号仓", "warehouse-c": "三号仓", "warehouse-d": "四号仓", "warehouse-e": "五号仓", "warehouse-f": "六号仓",
	"grade-a": "优", "grade-b": "良", "grade-c": "合格", "grade-d": "不合格", "grade-f": "失败", "unknown": "未知",
	"jan": "一月", "feb": "二月", "mar": "三月", "apr": "四月", "may": "五月", "jun": "六月", "jul": "七月", "aug": "八月", "sep": "九月", "oct": "十月", "nov": "十一月", "dec": "十二月",
	"inbound": "入库单", "quality": "质检结果", "settlement": "结算单", "supplier": "供应商", "permission": "权限", "session": "会话", "audit": "审计",
	"login": "登录", "logout": "登出", "create": "创建", "update": "更新", "delete": "删除", "read": "查看", "export": "导出", "grant": "授权", "revoke": "撤销",
	"admin": "管理员", "operator": "操作员", "viewer": "查看者", "owner": "所有者", "system": "系统", "manual": "人工", "automatic": "自动", "api": "接口",
	"csv": "CSV", "json": "JSON", "html": "HTML", "text": "文本", "success": "成功", "failure": "失败", "warning": "警告", "error": "错误",
}

func Label(k string) string {
	if v, ok := statusLabels[k]; ok {
		return v
	}
	return k
}
func Known(k string) bool { _, ok := statusLabels[k]; return ok }
func AllLabels() map[string]string {
	out := map[string]string{}
	for k, v := range statusLabels {
		out[k] = v
	}
	return out
}
