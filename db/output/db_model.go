package model


type Lb_change_log struct{
	Id	int64	//自增主键
	User_id	string	//用户ID
	Created_time	string	//创建时间
	Updated_time	string	//修改时间
}


type Load_balancer_pool struct{
	Id	int32	//自增主键
	Lb_ip	string	//服务器IP，内网IP，用于远程控制用
	Partition	int32	//分区ID，同样的分区ID代表一组负载均衡资源池
	Lb_type	int32	//lb服务器类型,1:lvs 2:nginx
	Created_time	string	//创建时间
	Updated_time	string	//更新时间
}


type Config struct{
	Id	int32	//自增id
	Key	string	//key
	Value	string	//value
	Created_time	string	//创建时间
	Updated_time	string	//修改时间
}


type Instance struct{
	Id	string	//负载均衡实例ID
	Name	string	//负载均衡实例名
	Detail	string	//负载均衡实例描述
	User_id	string	//用户ID
	Vip	string	//负载均衡虚IP
	Type	int32	//负载均衡类型：1：公网负载均衡,2:私网负载均衡
	Partition_id	int32	//分区ID
	Status	int32	//实例运行状态,1:已启用,2:已停用
	Max_conns	int32	//最大连接数
	Bandwidth	int32	//公网带宽
	Created_time	string	//创建时间
	Updated_time	string	//更新时间
}


type Instance_rule struct{
	Id	int64	//自增主键
	Instance_id	string	//关联instance表的外键
	Vip	string	//负载均衡ip，冗余字段
	L_port	int32	//负载均衡监听端口
	Level	int32	//转发类型:4:4层转发(tcp),7:7层转发(http)
	Rule	int32	//转发规则:1:轮询,2:最小连接数
	Rule_name	string	//规则名称
	Created_time	string	//创建时间
	Updated_time	string	//修改时间
}


type Instance_server struct{
	Id	int64	//自增主键
	Instance_id	string	//instance表外键
	Rule_id	int64	//instance_rule外键
	Router_ip	string	//宿主机路由器IP
	Router_port	int32	//宿主机路由器映射端口
	C_ip	string	//容器IP
	C_port	int32	//容器服务端口
	Created_time	string	//创建时间
	Updated_time	string	//修改时间
}


