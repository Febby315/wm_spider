const config={
    db_server_url:"http://localhost:5000",
    //1 分页模板解析
    db_spider_config_query:"/spiderdb/spiderconfig/list",//配置(查询)
    db_spider_list_source_add:"/spiderdb/listsource/add",//列表源(增加)
    //2 列表内容抓取
    db_list_source_query:"/spiderdb/listsource/list",//列表源(查询)
    db_list_source_edit:"/spiderdb/listsource/edit",//列表源(编辑)
    //3 列表价值清洗
    db_list_source_query:"/spiderdb/listsource/list",//列表源(查询)//要处理的数据表
    db_list_source_edit:"/spiderdb/listsource/edit",//列表源(编辑)//处理后的数据表
    db_list_detail_source_add:"/spiderdb/listdetailsource/add",//详情源(增加)//要处理的数据表
    //4 详情内容抓取
    db_list_detail_source_query:"/spiderdb/listdetailsource/list",//详情源(查询)
    db_list_detail_source_edit:"/spiderdb/listdetailsource/edit",//详情源(编辑)
    //5 详情价值清洗
    db_list_detail_source_query:"/spiderdb/listdetailsource/list",//详情源(查询)
    db_list_detail_source_edit:"/spiderdb/listdetailsource/edit",//详情源(编辑)
    db_list_detail_cleand_add:"/spiderdb/listdetailcleaned/add",//已清洗详情(增加)
    //6 附件资源下载
    db_list_detail_cleaned_query:"/spiderdb/listdetailcleaned/list",//已清洗详情(查询)
    db_list_detail_cleaned_edit:"/spiderdb/listdetailcleaned/edit",//已清洗详情(编辑)
    downImg_add:"/spiderdb_img_download/downImg/add",//附件(增加)
    downImg_findOne:"/spiderdb_img_download/downImg/findOne",//附件(单一查询)
    db_server_down_url:"http://localhost:3001",//*附件下载服务地址
    download_do:"/spiderdb_img_download/updownload/download",//*附件下载接口
    //7 附件资源清洗
    db_list_detail_cleaned_query:"/spiderdb/listdetailcleaned/list",//已清洗详情(查询)
    db_list_detail_cleaned_edit:"/spiderdb/listdetailcleaned/edit",//已清洗详情(编辑)
    //8 数据提纯抽离
    db_list_detail_cleaned_query:"/spiderdb/listdetailcleaned/list",//已清洗详情(查询)
    db_list_detail_cleaned_edit:"/spiderdb/listdetailcleaned/edit",//已清洗详情(编辑)
    db_article_add:"/spiderdb/article/add",//文章(增加)
    db_recruit_add:"/spiderdb/recruit/add",//招聘(增加)
    db_sell_add:"/spiderdb/sell/add",//供应(增加)
    db_buy_add:"/spiderdb/buy/add",//求购(增加)
    //9 抽注文章标签
    db_article_query:"/spiderdb/article/list",//文章(查询)
    db_article_edit:"/spiderdb/article/edit",//文章(编辑)
    db_tagArticle_add:"/spiderdb/tagArticle/add",//标签文章(增加)
}

//复制对象
function clone(myObj){
    var myNewObj = new Object(); 
    if(typeof myObj == 'object'){
        for(var i in myObj){ i!="_id"?myNewObj[i] = clone(myObj[i]):null; }
    }else{ myNewObj = myObj; }
    return myNewObj; 
}

module.exports = config;
module.exports.clone = clone;