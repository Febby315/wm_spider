const config={
	db_server_url:"http://localhost:5000",
	db_buy_add:"/spiderdb/buy/add",//要处理的数据表
   	db_list_detail_cleaned_query:"/spiderdb/listdetailcleaned/list",//要处理的数据表
	db_list_detail_cleaned_edit:"/spiderdb/listdetailcleaned/edit"//处理后的数据表
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