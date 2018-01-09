const config={
	db_server_url:"http://localhost:5000",
	db_list_detail_source_add:"/spiderdb/listdetailsource/add",//要处理的数据表
   	db_list_source_query:"/spiderdb/listsource/list",//要处理的数据表
	db_list_source_edit:"/spiderdb/listsource/edit"//处理后的数据表
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