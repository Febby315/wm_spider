const config={
	db_server_url:"http://localhost:5000",
    db_spider_config_query:"/spiderdb/spiderconfig/list",//获取配置表信息
	db_spider_list_source_add:"/spiderdb/listsource/add"
}

//复制对象
function clone(myObj){ 
	if(typeof(myObj) != 'object') return myObj; 
	if(myObj == null) return myObj; 
	var myNewObj = new Object(); 
	for(var i in myObj) 
	{
		if(i!=='_id')
		myNewObj[i] = clone(myObj[i]); 
	}
 return myNewObj; 
}

module.exports = config;
module.exports.clone= clone;