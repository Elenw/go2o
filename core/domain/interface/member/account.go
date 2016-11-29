/**
 * Copyright 2015 @ z3q.net.
 * name : account
 * author : jarryliu
 * date : 2015-07-24 08:48
 * description :
 * history :
 */
package member

const (
	// 余额账户
	AccountBalance = 1
	// 积分账户
	AccountIntegral = 2
	// 赠送账户
	AccountPresent = 3
	// 流通金账户
	AccountFlow = 4
)

const (
	// 用户充值
	ChargeByUser int32 = 1
	// 系统自动充值
	ChargeBySystem int32 = 2
	// 客服充值
	ChargeByService int32 = 3
	// 退款充值
	ChargeByRefund int32 = 4
)

const (
	// 会员充值
	KindBalanceCharge int32 = 1
	// 系统充值
	KindBalanceSystemCharge int32 = 2
	// 支付抵扣
	KindBalanceDiscount int32 = 3
	// 退款
	KindBalanceRefund int32 = 4
	// 转入
	KindBalanceTransferIn int32 = 5
	// 转出
	KindBalanceTransferOut int32 = 6
	// 失效
	KindBalanceExpired int32 = 7
	// 冻结
	KindBalanceFreeze int32 = 8
	// 解冻
	KindBalanceUnfreeze int32 = 9

	// 客服充值
	KindBalanceServiceCharge int32 = 15
	// 客服扣减
	KindBalanceServiceDiscount int32 = 16
)

const (
	// 赠送金额
	KindPresentAdd int32 = 1
	// 抵扣奖金
	KindPresentDiscount int32 = 2
	// 转入
	KindPresentTransferIn int32 = 5
	// 转出
	KindPresentTransferOut int32 = 6
	// 失效
	KindPresentExpired int32 = 7
	// 冻结
	KindPresentFreeze int32 = 8
	// 解冻
	KindPresentUnfreeze int32 = 9
	// 提现到余额
	KindPresentTakeOutToBalance int32 = 11
	// 提现到银行卡(人工提现)
	KindPresentTakeOutToBankCard int32 = 12
	// 提现到第三方
	KindPresentTakeOutToThirdPart int32 = 13
	// 提现退还到银行卡
	KindPresentTakeOutRefund int32 = 14
	// 支付单退款
	KindPresentPaymentRefund int32 = 15

	// 客服赠送
	KindPresentServiceAdd int32 = 21
	// 客服扣减
	KindPresentServiceDiscount int32 = 22
)

const (
	KindGrow = 7 // 增利

	//KindCommission = 9 // 手续费

	// 赠送
	//KindBalancePresent = 3

	// 流通账户
	KindBalanceFlow int32 = 4 // 账户流通

	// 提现
	//KindBalanceApplyCash = 11
	// 转账
	KindBalanceTransfer int32 = 12
	StatusOK                  = 1
)

const (
	// 赠送
	TypeIntegralPresent = 1
	// 积分抵扣
	TypeIntegralDiscount = 2
	// 积分冻结
	TypeIntegralFreeze = 3
	// 积分解冻
	TypeIntegralUnfreeze = 4
	// 购物赠送
	TypeIntegralShoppingPresent = 5
	// 支付抵扣
	TypeIntegralPaymentDiscount = 6
)

type (
	IAccount interface {
		// 获取领域对象编号
		GetDomainId() int32

		// 获取账户值
		GetValue() *Account

		// 保存
		Save() (int32, error)

		// 设置优先(默认)支付方式, account 为账户类型
		SetPriorityPay(account int, enabled bool) error

		// 根据编号获取余额变动信息
		GetBalanceInfo(id int32) *BalanceInfo

		// 根据号码获取余额变动信息
		// GetBalanceInfoByNo(no string) *BalanceInfo

		// 保存余额变动信息
		SaveBalanceInfo(*BalanceInfo) (int32, error)

		// 获取赠送账户日志
		GetPresentLog(id int32) *PresentLog

		// 充值,客服操作时,需提供操作人(relateUser)
		ChargeForBalance(chargeType int32, title string, outerNo string, amount float32, relateUser int32) error

		// 扣减余额
		DiscountBalance(title string, outerNo string, amount float32, relateUser int32) error

		// 冻结余额
		Freeze(title string, outerNo string, amount float32, relateUser int32) error

		// 解冻金额
		Unfreeze(title string, outerNo string, amount float32, relateUser int32) error

		// 赠送金额,客服操作时,需提供操作人(relateUser)
		ChargeForPresent(title string, outerNo string, amount float32, relateUser int32) error

		// 赠送金额(指定业务类型)
		ChargePresentByKind(kind int32, title string, outerNo string, amount float32, relateUser int32) error

		// 扣减奖金,mustLargeZero是否必须大于0, 赠送金额存在扣为负数的情况
		DiscountPresent(title string, outerNo string, amount float32,
			relateUser int32, mustLargeZero bool) error

		// 冻结赠送金额
		FreezePresent(title string, outerNo string, amount float32, relateUser int32) error

		// 解冻赠送金额
		UnfreezePresent(title string, outerNo string, amount float32, relateUser int32) error

		// 流通账户余额变动，如扣除,amount传入负数金额
		ChargeFlowBalance(title string, tradeNo string, amount float32) error

		// 支付单抵扣消费,tradeNo为支付单单号
		PaymentDiscount(tradeNo string, amount float32, remark string) error

		//　增加积分
		AddIntegral(iType int, outerNo string, value int, remark string) error

		// 积分抵扣
		IntegralDiscount(logType int, outerNo string, value int, remark string) error

		// 冻结积分,当new为true不扣除积分,反之扣除积分
		FreezesIntegral(value int, new bool, remark string) error

		// 解冻积分
		UnfreezesIntegral(value int, remark string) error

		// 退款
		RequestBackBalance(backType int, title string, amount float32) error

		// 完成退款
		FinishBackBalance(id int32, tradeNo string) error

		// 申请提现,applyType：提现方式,返回info_id,交易号 及错误
		RequestTakeOut(takeKind int32, title string, amount float32, commission float32) (int32, string, error)

		// 确认提现
		ConfirmTakeOut(id int32, pass bool, remark string) error

		// 完成提现
		FinishTakeOut(id int32, tradeNo string) error

		// 将冻结金额标记为失效
		FreezeExpired(accountKind int, amount float32, remark string) error

		// 转账
		TransferAccounts(accountKind int, toMember int32, amount float32,
			csnRate float32, remark string) error

		// 接收转账
		ReceiveTransfer(accountKind int, fromMember int32, tradeNo string,
			amount float32, remark string) error

		// 转账余额到其他账户
		TransferBalance(kind int32, amount float32, tradeNo string, toTitle, fromTitle string) error

		// 转账返利账户,kind为转账类型，如 KindBalanceTransfer等
		// commission手续费
		TransferPresent(kind int32, amount float32, commission float32, tradeNo string,
			toTitle string, fromTitle string) error

		// 转账活动账户,kind为转账类型，如 KindBalanceTransfer等
		// commission手续费
		TransferFlow(kind int32, amount float32, commission float32, tradeNo string,
			toTitle string, fromTitle string) error

		// 将活动金转给其他人
		TransferFlowTo(memberId int32, kind int32, amount float32, commission float32,
			tradeNo string, toTitle string, fromTitle string) error
	}

	// 余额变动信息
	BalanceInfo struct {
		Id       int32  `db:"id" auto:"yes" pk:"yes"`
		MemberId int32  `db:"member_id"`
		TradeNo  string `db:"trade_no"`
		Kind     int32  `db:"kind"`
		Type     int    `db:"type"`
		Title    string `db:"title"`
		// 金额
		Amount float32 `db:"amount"`
		// 手续费
		CsnAmount float32 `db:"csn_amount"`
		// 引用编号
		RefId      int32 `db:"ref_id"`
		State      int   `db:"state"`
		CreateTime int64 `db:"create_time"`
		UpdateTime int64 `db:"update_time"`
	}

	// 账户值对象
	Account struct {
		// 会员编号
		MemberId int32 `db:"member_id" pk:"yes" json:"memberId"`
		// 积分
		Integral int `db:"integral"`
		// 不可用积分
		FreezeIntegral int `db:"freeze_integral"`
		// 余额
		Balance float32 `db:"balance" json:"balance"`
		// 不可用余额
		FreezeBalance float32 `db:"freeze_balance" json:"freezesFee"`
		// 失效的账户余额
		ExpiredBalance float32 `db:"expired_balance"`
		//奖金账户余额
		PresentBalance float32 `db:"present_balance" json:"presentBalance"`
		//冻结赠送金额
		FreezePresent float32 `db:"freeze_present" json:"FreezePresent"`
		//失效的赠送金额
		ExpiredPresent float32 `db:"expired_present"`
		//总赠送金额
		TotalPresentFee float32 `db:"total_present_fee" json:"totalPresentFee"`
		//流动账户余额
		FlowBalance float32 `db:"flow_balance" json:"flowBalance"`
		//当前理财账户余额
		GrowBalance float32 `db:"grow_balance" json:"growBalance"`
		//理财总投资金额,不含收益
		GrowAmount float32 `db:"grow_amount" json:"growAmount"`
		//当前收益金额
		GrowEarnings float32 `db:"grow_earnings" json:"growEarnings"`
		//累积收益金额
		GrowTotalEarnings float32 `db:"grow_total_earnings" json:"growTotalEarnings"`
		//总消费金额
		TotalConsumption float32 `db:"total_consumption" json:"totalFee"`
		//总充值金额
		TotalCharge float32 `db:"total_charge" json:"totalCharge"`
		//总支付额
		TotalPay float32 `db:"total_pay" json:"totalPay"`
		// 优先(默认)支付选项
		PriorityPay int `db:"priority_pay"`
		//更新时间
		UpdateTime int64 `db:"update_time" json:"updateTime"`
	}

	// 积分记录
	IntegralLog struct {
		// 编号
		Id int32 `db:"id" pk:"yes" auto:"yes"`
		// 会员编号
		MemberId int32 `db:"member_id"`
		// 类型
		Type int `db:"type"`
		// 关联的编号
		OuterNo string `db:"outer_no"`
		// 积分值
		Value int `db:"value"`
		// 备注
		Remark string `db:"remark"`
		// 创建时间
		CreateTime int64 `db:"create_time"`
	}

	// 余额日志
	BalanceLog struct {
		Id       int32  `db:"id" auto:"yes" pk:"yes"`
		MemberId int32  `db:"member_id"`
		OuterNo  string `db:"outer_no"`
		// 业务类型
		BusinessKind int32 `db:"kind"`

		Title string `db:"title"`
		// 金额
		Amount float32 `db:"amount"`
		// 手续费
		CsnFee float32 `db:"csn_fee"`
		// 关联操作人,仅在客服操作时,记录操作人
		RelateUser int32 `db:"rel_user"`
		// 状态
		State int32 `db:"state"`
		// 备注
		Remark string `db:"remark"`
		// 创建时间
		CreateTime int64 `db:"create_time"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
	}

	// 赠送账户日志
	PresentLog struct {
		Id int32 `db:"id" auto:"yes" pk:"yes"`
		// 会员编号
		MemberId int32 `db:"member_id"`
		// 外部单号
		OuterNo string `db:"outer_no"`
		// 业务类型
		BusinessKind int32 `db:"kind"`
		// 标题
		Title string `db:"title"`
		// 金额
		Amount float32 `db:"amount"`
		// 手续费
		CsnFee float32 `db:"csn_fee"`
		// 关联操作人,仅在客服操作时,记录操作人
		RelateUser int32 `db:"rel_user"`
		// 状态
		State int32 `db:"state"`
		// 备注
		Remark string `db:"remark"`
		// 创建时间
		CreateTime int64 `db:"create_time"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
	}
)
