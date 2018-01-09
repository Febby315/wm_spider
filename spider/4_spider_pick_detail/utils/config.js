const config={
	db_server_url:"http://localhost:5000",
	db_list_detail_source_query:"/spiderdb/listdetailsource/list",
	db_list_detail_source_edit:"/spiderdb/listdetailsource/edit",
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
module.exports.clone= clone;
module.exports = config;
