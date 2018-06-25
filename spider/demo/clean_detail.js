// 详情清洗脚本DEMO文件
if(/www\.moa\.gov\.cn\/fwllm\/qgxxlb\/qg\/\d+\/t\d+_\d+\.htm/gm.test(param.detail_url)){
	$=cheerio.load($("div.zleft").eq(0).html(),{decodeEntities: false});
}else{ throw "该解析模板无法解析此链接:"+param.detail_url }

var outResult=param.table_info;
var regSource=/来源：(.*)/gm.exec($(".hui_12-12").text());
regSource?outResult.source=regSource[1]:null;
param.pub_time=moment($(".hui_12-12").text().replace(/作者(.*)/,""),"YYYY-MM-dd HH:mm").format("YYYY-MM-DD HH:mm:ss");

$=cheerio.load($("#TRS_AUTOADD").eq(0).html(),{decodeEntities: false});
var attachFilesSour=[];    //文章附件
$("a").each(function (i, ele) {
    var name=$(ele).text();
    var src=url.resolve(param.detail_url,$(ele).attr("href"));
    attachFilesSour.push({ 'name':name ,'sour_url':src });
});

function clean(cleanHtml){  //清理函数(支持css选择器)
    //标签内容置空
    cleanHtml.content.forEach(function (selector,index) {
        $(selector).html(null);
    });
    //移除标签
    cleanHtml.tag.forEach(function (selector,index) {
        $(selector).remove();
    });
    //移除属性
    cleanHtml.param.forEach(function (obj,index) {
        obj.params.forEach(function (p,i) {
            $(obj.selector).removeAttr(p);
        });
    });
    //替换标签(仅支持双标签)
    cleanHtml.replaceWith.forEach(function (obj,index) {
        $(obj.selector).each(function(){
            $(this).replaceWith('<'+obj.tag+'>'+$(this).html()+'</'+obj.tag+'>');
        });
    });
}
var cleanHtml={};//清理规则
cleanHtml.content=[];
cleanHtml.tag=["link,script,style,iframe"];
cleanHtml.param=[
	{selector:"*",params:["id","class","style","alt"]}
];
cleanHtml.replaceWith=[
	{selector:"section",tag:"p"},
	{selector:"a",tag:"span"}
];
clean(cleanHtml);//执行清洗

var contentImageSour=[];    //内容图片
$("img").each(function (i, element) {
    var src=url.resolve(param.detail_url,$(element).attr('src').trim());
    contentImageSour.push({ 'sour_url': src });
    $(element).replaceWith('<img src="' + src + '"/>');
});

param.contentImageSour=JSON.stringify(contentImageSour);//文章图片
param.attachFilesSour=JSON.stringify(attachFilesSour);//文章附件
param.summary=$.text().replace(/(&nbsp;|\n|\s)+/gm," ").trim().substr(0,100);//文章摘要
param.content_sour=$.html();//文章内容