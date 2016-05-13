function mod_selection(){
		
	if(document.selection){ //ie
		return document.selection.createRange();
	}
	else{  //标准
		return window.getSelection();
	}
	
}

function insertText(id, text)
{
	$('#'+id).focus();
    var str = mod_selection();
    console.log(str);
	str.text = text;
}


function show_url_type(obj) {
	var num = 4;
	for (var i=1; i<=num; i++){
		if (obj==i){ 
			$('#url_type_'+i).show();
		} else {
			$('#url_type_'+i).hide();
		}
	}
}

function show_url() {
	var type = $("input[type='radio'][name='data[sourcetype]']:checked").val();
	art.dialog({id:'test_url',iframe:'?m=collection&c=node&a=public_url&sourcetype='+type+'&urlpage='+$('#urlpage_'+type).val()+'&pagesize_start='+$("input[name='data[pagesize_start]']").val()+'&pagesize_end='+$("input[name='data[pagesize_end]']").val()+'&par_num='+$("input[name='data[par_num]']").val(), title:'测试地址', width:'700', height:'450'}, '', function(){art.dialog({id:'test_url'}).close()});			
}


function show_div(id) {
	for (var i=1;i<=4;i++) {
		if (id==i) {
			$('#tab_'+i).addClass('on');
			$('#show_div_'+i).show();
		} else {
			$('#tab_'+i).removeClass('on');
			$('#show_div_'+i).hide();
		}
	}
}

function show_nextpage(value) {
	if (value == 2) {
		$('#nextpage').show();
	} else {
		$('#nextpage').hide();
	}
}

function anti_selectall(obj) {
	$("input[name='"+obj+"']").each(function(i,n){
		if (this.checked) {
			this.checked = false;
		} else {
			this.checked = true;
		}});
}

function selectall(name) {
	if ($("#check_box").attr("checked")=='checked') {
		$("input[name='"+name+"']").each(function() {
  			$(this).attr("checked","checked");
			
		});
	} else {
		$("input[name='"+name+"']").each(function() {
  			$(this).removeAttr("checked");
		});
	}
}

var i =0;
function add_caiji() {
	var html = '<tbody id="customize_config_'+i+'"><tr style="background-color:#FBFFE4"><td>规则名：</td><td><input type="text" name="customize_config[name][]" class="input-text" /></td><td>规则英文名：</td><td><input type="text" name="customize_config[en_name][]" class="input-text" /></td></tr><tr><td width="120">匹配规则：</td><td><textarea rows="5" cols="40" name="customize_config[rule][]" id="rule_'+i+'"></textarea> <br/>使用"<a href="javascript:insertText(\'rule_'+i+'\', \'[内容]\')">[内容]</a>"作为通配符</td><td width="120">过滤选项：</td><td><textarea rows="5" cols="50" name="customize_config[html_rule][]" id="content_html_rule_'+i+'"></textarea><input type="button" value="选择" class="button"  onclick="html_role(\'content_html_rule_'+i+'\', 1)"></td></tr></tbody>';
	$('#customize_config').append(html);
	i++;
}


function html_role(id, type) {
	art.dialog({
		id:'test_url',
		content:'<label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_1" value="&lt;p([^&gt;]*)&gt;(.*)&lt;/p&gt;[|]"> &lt;p&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_2" value="&lt;a([^&gt;]*)&gt;(.*)&lt;/a&gt;[|]"> &lt;a&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_3" value="&lt;script([^&gt;]*)&gt;(.*)&lt;/script&gt;[|]"> &lt;script&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_4" value="&lt;iframe([^&gt;]*)&gt;(.*)&lt;/iframe&gt;[|]"> &lt;iframe&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_5" value="&lt;table([^&gt;]*)&gt;(.*)&lt;/table&gt;[|]"> &lt;table&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_6" value="&lt;span([^&gt;]*)&gt;(.*)&lt;/span&gt;[|]"> &lt;span&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_7" value="&lt;b([^&gt;]*)&gt;(.*)&lt;/b&gt;[|]"> &lt;b&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_8" value="&lt;img([^&gt;]*)&gt;[|]"> &lt;img&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_9" value="&lt;object([^&gt;]*)&gt;(.*)&lt;/object&gt;[|]"> &lt;object&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_10"  value="&lt;embed([^&gt;]*)&gt;(.*)&lt;/embed&gt;[|]"> &lt;embed&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_11"  value="&lt;param([^&gt;]*)&gt;(.*)&lt;/param&gt;[|]"> &lt;param&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_12"  value="&lt;div([^&gt;]*)&gt;[|]"> &lt;div&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_13"  value="&lt;/div&gt;[|]"> &lt;/div&gt;</label><label class="ib" style="width:120px"><input type="checkbox" name="html_rule" id="_14"  value="&lt;!--([^&gt;]*)--&gt;[|]"> &lt;!-- --&gt;</label><br/><div class="bk15"></div><center><input type="button" value="全选" class="button" onclick="selectall(\'html_rule\')"><input type="button" class="button"  value="反选" onclick="anti_selectall(\'html_rule\')"></center>',
		width:'500',
		height:'150',
		lock: false}, function(){
			var old = $("textarea[name='"+id+"']").val();
			var str = '';
			
			$("input[name='html_rule']:checked").each(function(){
				str+=$(this).val()+"\n";
			});

			$((type == 1 ? "#"+id :"textarea[name='"+id+"']")).val((old ? old+"\n" : '')+str);
		}, function(){
			art.dialog({id:'test_url'}).close()
		});
}

function test_spider(id) {	
//	window.open('public_test/nodeid/'+id, 'test', 'width=700,height=500,top=0,left=0,scrollbars=1');

	art.dialog({id:'test'}).close();
	art.dialog({
		title:'测试采集',id:'test',iframe:'public_test/nodeid/'+id,width:'700',height:'500'
		}, '', function(){
			art.dialog({id:'test'}).close()
			});

}

function show_content(url, nodeid) {
	art.dialog({id:'test',content:'加载中...',width:'100',height:'30', lock:true});
	$.get("admin.php?m=Collection&a=public_test_content&nodeid="+nodeid+"&url="+url, function(data){
		art.dialog({id:'test'}).close();
		art.dialog({
			title:'内容查看',id:'test',content:'<textarea rows="26" cols="90">'+data+'</textarea>',width:'500',height:'400', lock:false
		});
	});
}

function selectall(name) {
	if ($("#check_box").attr("checked")=='checked') {
		$("input[name='"+name+"']").each(function() {
  			$(this).attr("checked","checked");
			
		});
	} else {
		$("input[name='"+name+"']").each(function() {
  			$(this).removeAttr("checked");
		});
	}
}

function check_checkbox(){
	res = false;
	input = $("input[type='checkbox']:checked");

	if(input.length> 0){
		if(confirm('您确定要删除吗？')){
			res = true;
		}
	}else{
		alert('还没有选择数据！');
	}
	return res;
}

function re_url(delHis) {
	input = null;
	if(delHis){
		$('#history').val('1');
	}
	return false;
}

function copy_spider(id) {
	art.dialog({id:'test'}).close();
	art.dialog({title:'复制采集',id:'test',iframe:'admin.php?m=Collection&a=copy&nodeid='+id,width:'420',height:'120'}, function(){var d = art.dialog({id:'test'}).data.iframe;var form = d.document.getElementById('dosubmit');form.click();return false;}, function(){art.dialog({id:'test'}).close()});
}
