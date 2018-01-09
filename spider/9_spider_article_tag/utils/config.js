const config={
	db_server_url:"http://localhost:5000",
	db_article_query:"/spiderdb/article/list",//获取清理后的数据表信息
	db_article_edit:"/spiderdb/article/edit",//处理清理后的数据表
	db_tagArticle_add:"/spiderdb/tagArticle/add"//
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