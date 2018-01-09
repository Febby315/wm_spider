var cheerio = require('cheerio');//jquery 方式操作
var send = require('./utils/send');
var config = require('./utils/config');
var moment = require('moment');
//查找配置文件
var result = send.post(config.db_server_url + config.db_article_query, { conf:{dealStatus: 0}});
var data = result.data;
for (var i in data) {
	var param = config.clone(data[i]);
	param.keyStatus = '0';
	param.nlpDate = moment().format('YYYY-MM-DD HH:mm:ss');
	param.keyWordList = '[' + param.articleNavTitle + ',' + param.articleClassification + ']';
	param.keyTitleList = '[' + param.articleNavTitle + ']';//1
    let editparam = { _id: data[i]._id, version: data[i].version, dealStatus: 10 };
	if(param.info_flag==="sxnyw_trsq"){
		var privateField={};
		var content=param.articleContent.replace(/\s+/gm,"").replace(/（[\d\.]+%-[\d\.]+%）/gm,"");
		var reg=/据.+(渭北).(关中).(陕南).(陕北)?.+（(0-20cm)）.+平均值分别为([\d\.]+)%?.([\d\.]+)%?.([\d\.]+)%?.(([\d\.]+)%?.)?（(20-40cm)）分别为([\d\.]+)%?.([\d\.]+)%?.([\d\.]+)%?.(([\d\.]+)%.)?/gm.exec(content);
		
		var soilMoisture={};
		if(reg){
			soilMoisture.zone=[]; soilMoisture.cm0_20=[]; soilMoisture.cm20_40=[];soilMoisture.unit="%";
			soilMoisture.date=moment(param.articleCreateDateTime).subtract(2,"days").format("YYYY-MM-DD");
			for(var n=1;n<5;n++){
				if(reg[n]){
					if(n===4){
						soilMoisture.zone.push(reg[n]); soilMoisture.cm0_20.push(parseFloat(reg[n+5+1])); soilMoisture.cm20_40.push(parseFloat(reg[n+11+1]));
					}else{
						soilMoisture.zone.push(reg[n]); soilMoisture.cm0_20.push(parseFloat(reg[n+5]));  soilMoisture.cm20_40.push(parseFloat(reg[n+11]));
					}
				}
			}
		}
		privateField.soilMoisture=soilMoisture;
		param.privateField=JSON.stringify(privateField,null,"\t");
	}
	// console.log('param.keyTitleList',param.keyTitleList);
	param.article_id = data[i]._id;
	var result = send.post(config.db_server_url + config.db_tagArticle_add, param);
	//处理完的数据修改状态
	send.post(config.db_server_url + config.db_article_edit, editparam);
}