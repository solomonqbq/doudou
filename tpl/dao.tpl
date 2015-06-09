package {{ package }}
const{
    QUERY_{{ m. }}
}

{% macro model(m) %}
    {% for ip in m.Rs_IPS %}
    real_server {{ ip }} {{ m.Vs_Port }} {
            weight 1
            TCP_CHECK {
                    connect_timeout 3
                    nb_get_retry 3
                    delay_before_retry 3
                    connect_port {{ m.Vs_Port }}
            }
        }
    {% endfor %}

type Instance_server struct{
	Instance_id	string	//instance表外键
	Rule_id	int64	//instance_rule外键
	Host_ip	string	//宿主机出口IP
	Host_port	int32	//宿主机映射端口
	C_ip	string	//容器IP
	C_port	int32	//容器服务端口
	Created_time	string	//创建时间
	Updated_time	string	//修改时间
	Deleted	int32	//删除标志位,0:未删除 1:已删除
	Changed	int32	//是否修改过标志 0:未修改 1:已修改
}

virtual_server {{ m.Vs_IP }} {{ m.Vs_Port }} {
    delay_loop 6
    lb_algo rr
    lb_kind DR
    persistence_timeout 50
    protocol TCP
    {% for ip in m.Rs_IPS %}
    real_server {{ ip }} {{ m.Vs_Port }} {
            weight 1
            TCP_CHECK {
                    connect_timeout 3
                    nb_get_retry 3
                    delay_before_retry 3
                    connect_port {{ m.Vs_Port }}
            }
        }
    {% endfor %}
}
{% endmacro %}

{% for m in ms %}
    {{ vs(m) }}
{% endfor %}

