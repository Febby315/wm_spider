# Mongo同步查询接口说明文档

## 1. etl_mongo_src介绍

>本程序设计初衷为:
>为mongo数据库中的结果数据提供统一查询接口
>并为同步中间件(etl_do_*)提供查询

    声明:本程序接口均已允许跨域请求

## 2. 接口说明

### /spiderdb/sysappflag/findone

*POST [/spiderdb/sysappflag/findone](http://localhost:5002/spiderdb/sysappflag/findone)* **允许检索查询接口**

```javascript
{
    "conf":{    //查询条件
        "appid":"sxny",             //程序标识
        "passwd":"sxny"             //访问密码
    }
}
```

    声明:conf中允许包含其他属性,其将作为查询条件(完全匹配查询)
    appid:该参数为程序标识
    passwd:该参数为访问密码

### /spiderdb/sysdatapoint/findone

*POST [/spiderdb/sysdatapoint/findone](http://localhost:5002/spiderdb/sysdatapoint/findone)* **同步记录查询接口**

```javascript
{
    "conf":{    //查询条件
        "appid":"sxny",             //程序标识
        "passwd":"sxny",            //访问密码
        "tablename":"tagArticle"    //需要查询同步进度的表名
    }
}
```

    声明:conf中允许包含其他属性,其将作为查询条件(完全匹配查询)
    appid:该参数为程序标识
    passwd:该参数为访问密码
    tablename:允许值"tagArticle","sellTransaction","buyTransaction","price","recruit","downImg"

### /spiderdb/sysdatapoint/edit

*POST [/spiderdb/sysdatapoint/edit](http://localhost:5002/spiderdb/sysdatapoint/edit)* **同步记录修改接口**

```javascript
{
    "conf":{    //查询条件
        "appid":"sxny",             //程序标识
        "passwd":"sxny",            //访问密码
        "tablename":"tagArticle"    //需要查询同步进度的表名
        "last_sys_id":""    //最新进度ID标识
    }
}
```

    声明:conf中允许包含其他属性,其将作为查询条件(完全匹配查询)
    appid:该参数为程序标识
    passwd:该参数为访问密码
    tablename:允许值"tagArticle","sellTransaction","buyTransaction","price","recruit","downImg"
    last_sys_id:该值将用于记录您最后一次同步到哪一条数据,方便下次从此位置继续同步数据

### /spiderdb/tagArticle/list

*POST [/spiderdb/tagArticle/list](http://localhost:5002/spiderdb/tagArticle/list)* **文章信息查询接口**

```javascript
{
    "conf":{    //查询条件
        "info_flag": "sxny_yanan",          //"_all"、"sxny"、"sxny,sxny_yanan"
        "_id":"5ad44b5c7cb5d81e2c6ff30f"    //从该条数据开始查询{pageSize}条记录
    },
    "pageSize":10                           //查询条数
}
```

    默认值:info_flag:"_all",pageSize:10
    声明:conf中允许包含其他属性,其将作为查询条件(完全匹配查询)
    _id:为空时将返回首条记录的_id
    info_flag:多条记录可使用','分隔,如'sxny,sxny_yanan'

### /spiderdb/sell/list

*POST [/spiderdb/sell/list](http://localhost:5002/spiderdb/sell/list)* **供应信息查询接口**

```javascript
{
    "conf":{    //查询条件
        "info_flag": "sxny_yanan",  //"_all"、"sxny"、"sxny,sxny_yanan"
        "_id":"5ad44b5c7cb5d81e2c6ff30f"    //从该条数据开始查询{pageSize}条记录
    },
    "pageSize":10   //查询条数
}
```

    默认值:info_flag:"_all",pageSize:10
    声明:conf中允许包含其他属性,其将作为查询条件(完全匹配查询)
    _id:为空时将返回首条记录的_id
    info_flag:多条记录可使用','分隔,如'sxny,sxny_yanan'

### /spiderdb/buy/list

*POST [/spiderdb/buy/list](http://localhost:5002/spiderdb/buy/list)* **求购信息查询接口**

```javascript
{
    "conf":{    //查询条件
        "info_flag": "sxny_yanan",  //"_all"、"sxny"、"sxny,sxny_yanan"
        "_id":"5ad44b5c7cb5d81e2c6ff30f"    //从该条数据开始查询{pageSize}条记录
    },
    "pageSize":10   //查询条数
}
```

    默认值:info_flag:"_all",pageSize:10
    声明:conf中允许包含其他属性,其将作为查询条件(完全匹配查询)
    _id:为空时将返回首条记录的_id
    info_flag:多条记录可使用','分隔,如'sxny,sxny_yanan'

### /spiderdb/price/list

*POST [/spiderdb/price/list](http://localhost:5002/spiderdb/price/list)* **价格信息查询接口**

```javascript
{
    "conf":{    //查询条件
        "info_flag": "sxny_yanan",  //"_all"、"sxny"、"sxny,sxny_yanan"
        "_id":"5ad44b5c7cb5d81e2c6ff30f"    //从该条数据开始查询{pageSize}条记录
    },
    "pageSize":10   //查询条数
}
```

    默认值:info_flag:"_all",pageSize:10
    声明:conf中允许包含其他属性,其将作为查询条件(完全匹配查询)
    _id:为空时将返回首条记录的_id
    info_flag:多条记录可使用','分隔,如'sxny,sxny_yanan'

### /spiderdb/recruit/list

*POST [/spiderdb/recruit/list](http://localhost:5002/spiderdb/recruit/list)* **招聘信息查询接口**

```javascript
{
    "conf":{    //查询条件
        "info_flag": "sxny_yanan",  //"_all"、"sxny"、"sxny,sxny_yanan"
        "_id":"5ad44b5c7cb5d81e2c6ff30f"    //从该条数据开始查询{pageSize}条记录
    },
    "pageSize":10   //查询条数
}
```

    默认值:info_flag:"_all",pageSize:10
    声明:conf中允许包含其他属性,其将作为查询条件(完全匹配查询)
    _id:为空时将返回首条记录的_id
    info_flag:多条记录可使用','分隔,如'sxny,sxny_yanan'

### /spiderdb/downImg/list

*POST [/spiderdb/downImg/list](http://localhost:5002/spiderdb/downImg/list)* **附件信息查询接口**

```javascript
{
    "conf":{    //查询条件
        "dealStatus": "0",  //"_all"、"0"、"sxny,sxny_yanan"
        "_id":"5ad44b5c7cb5d81e2c6ff30f"    //从该条数据开始查询{pageSize}条记录
    },
    "pageSize":10   //查询条数
}
```

    默认值:info_flag:"_all",pageSize:10
    声明:conf中允许包含其他属性,其将作为查询条件(完全匹配查询)
    _id:为空时将返回首条记录的_id
    dealStatus:本接口请忽略传递此参数