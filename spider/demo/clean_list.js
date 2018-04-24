$(".u-list>ul>li").each(function(index,ele){
	var href=$("a",ele).attr("href").toString().trim();
	var title=$("a",ele).text().trim();
	sendAdd(href,title);
});