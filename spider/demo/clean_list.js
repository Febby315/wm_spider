// 列表清洗脚本DEMO文件
$(".u-list>ul>li").each(function(index,ele){
	var href=$("a",ele).attr("href").toString().trim();
	var title=$("a",ele).text().trim();
	sendAdd(href,title);
});