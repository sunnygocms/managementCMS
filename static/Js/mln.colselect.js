/**
 *栏目树分级显示下拉菜单组件
 *作者：CandySunPlus 孙凤鸣
 *example:
 *var colObj2 = {"Items":[
 *				{"name":"菜单项目标题1","parent_id":"0","id":"1","value":"菜单项目标题1","fun":function(){}},	
 *				...
 *				]}
 *$("#div").mlnColsel(colObj,{
 *				title:"栏目下拉菜单",
 *				value:"-1",
 *				width:100
 *			});
 *参数：parent_id 栏目所属，0为一级
 *		id 栏目ID，0为一级
 *		fun 回调函数
 **/
 (function($){
	$.fn.mlnColsel=function(data,setting){
		var dataObj={"Items":[
			{"name":"mlnColsel","parent_id":"-1","id":"-1"}
		]};
		var settingObj={
			title:"请选择",
			value:"-1",
			width:100,
			justLeaf:true
		};
		
		settingObj=$.extend(settingObj,setting);
		dataObj=$.extend(dataObj,data);
		var $this=$(this);
		var $colselbox=$(document.createElement("a")).addClass("colselect").attr({"href":"javascript:;"});
		var $colseltext=$(document.createElement("span")).text(settingObj.title);
		var $coldrop=$(document.createElement("ul")).addClass("dropmenu");
		var selectInput = ($.browser.msie&&$.browser.version<8)?document.createElement("<input name="+$this.attr("id")+" />"):document.createElement("input");
 			selectInput.type="hidden";
 			selectInput.value=settingObj.value;
 			selectInput.setAttribute("name",$this.attr("id"));
		var ids=$this.attr("id");
		$this.onselectstart=function(){return false;};
		$this.onselect=function(){document.selection.empty()};
		$colselbox.append($colseltext);
		$this.addClass("colsel").append($colselbox).append($coldrop).append(selectInput);
		
		$(dataObj.Items).each(function(i,n){
			var $item=$("<li/>");
			if(n.id==settingObj.value){
				$colseltext.text(n.name);
			}
			if(n.parent_id==0){
				$coldrop.append($item);
				$item.html("<span>"+n.name+"</span>").attr({"values":n.id,"id":"col_"+ids+"_"+n.id});
			}else{
				if($("#col_"+ids+"_"+n.parent_id).find("ul").length<=0){
					$("#col_"+ids+"_"+n.parent_id).append("<ul class=\"dropmenu rootmenu\"></ul>");
					$("#col_"+ids+"_"+n.parent_id).find("ul:first").append($item);
					$item.html("<span>"+n.name+"</span>").attr({"values":n.id,"id":"col_"+ids+"_"+n.id});
				}else{
					$("#col_"+ids+"_"+n.parent_id).find("ul:first").append($item);
					$item.html("<span>"+n.name+"</span>").attr({"values":n.id,"id":"col_"+ids+"_"+n.id});
				}
			}			
		});
		
		$this.find("li").each(function(){
			$(this).click(function(event){
				if(settingObj.justLeaf&&$(this).children("ul").length){
					
				}
				else{
					$colselbox.children("span").text($(this).find("span:first").text());
					$(selectInput).val($(this).attr("values"));
					hideMenu();	
				}
				event.stopPropagation();
			});
			if($(this).find("ul").length>0){
				$(this).addClass("menuout");
				$(this).hover(function(){
						$(this).removeClass("menuout");
						$(this).addClass("menuhover");
						$(this).find("ul:first").show()
					},function(){
						$(this).removeClass("menuhover");
						$(this).addClass("menuout");
						$(this).find("ul").each(function(){
							$(this).hide();
					});
				});
			}else{
				$(this).addClass("norout");
				$(this).hover(function(){
					$(this).removeClass("norout");
					$(this).addClass("norhover");
				},function(){
					$(this).removeClass("norhover");
					$(this).addClass("norout");
				});
			}
		});
		
        if($.browser.msie&&$.browser.version<8){
            //$(this).css({"margin-top":"-9px","top":"auto"});
            var $ul = $this.find("ul:first");
            $ul.css({"margin-left":-$ul.width()-5,"margin-top":20});
        }
		function hideMenu(){
			$this.bOpen=0;
			$(".rootmenu").hide();
			$coldrop.hide();
			$(document).unbind("click",hideMenu);
		}
		function openMenu(){
			$coldrop.show();
			$this.bOpen=1;
		}
		$colselbox.click(function(event){
            $(this).blur();
			if($this.bOpen){
				hideMenu();
			}else{
				openMenu();
				$(document).bind("click",hideMenu);
			}
			event.stopPropagation();
		});
		$(".rootmenu").each(function(){
			if($.browser.msie&&$.browser.version<8){
				$(this).css({"margin-top":"-9px","top":"auto"});
				$(this).css({"left":$(this).width()+"px"});
			}else{
				$(this).css({"margin-top":"-25px","top":"auto"});
				$(this).css({"left":135+"px"});
			}
			
		});	
	}
 })(jQuery);