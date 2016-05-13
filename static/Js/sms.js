$(document).ready(function() {
	//alert($("#channel_id").val());
	if($("#channel_id").val()==-1){
		$("#bodyform").hide();
	}else{
		$("#bodyform").show();
	}
	$("#type_id").change(function(){
		$('#searchForm').get(0).submit();
	});
	$("#channel_id").change(function(){
		if($(this).val()==-1){
			$("#issue").val('');
			$("#bodyform").hide();
			//$('#searchForm').get(0).submit();
		}else{
			$("#bodyform").show();
		}
		$('#searchForm').get(0).submit();
	});
	$('#searchForm').submit(function() {
		if ($("#channel_id").val()==-1){
			return false;
		}
		
	});
	
	$("td div.input").on("click",function(){
        var $div = $(this);
        if(!$div.find("input").length){
            var oldValue = $div.text();
            var $input = $("<input>").css({
                    "border":"0 none","width":"100%",height:"100%","padding":0
                }).val(oldValue).appendTo($div.empty()).keyup(function(event){
                    if(event.keyCode==13){
                        $input.blur();
                        event.preventDefault();
                        event.stopPropagation();
                        return false;
                    }
                }).blur(function(){
                    var newValue = $input.val();
                    $div.text(newValue);
                    if(newValue!=oldValue){
                        //$div.parents("tr:last").children().eq(0).text('有变化！');
//                        var id = $.trim($div.parents("tr:first").children().eq(0).text());
                        $div.parents("td:first").next().text(newValue.length);
                        $div.next().val(newValue);
                        for(i=0;i<keyword.length;i++){
                        	if(newValue.indexOf(keyword[i])>-1){
                        		console.log(keyword[i]);
                        		alert("“"+keyword[i]+"”是敏感词，请修改");
                        		break;
                        	}
                        }
                    }
                }).focus();
        }
    }).each(function(){
    	var $this = $(this);
    	if($this.text().length>0){
    		$this.parent().next().text($this.text().length);
    	}
    	
    });
});
function doSubmit(){
	var returnValue = true;
	$("[name='content[]']").each(function(){
		var val = $(this).val(),
			l = keyword.length;
		if(val.length>0){
			for(i=0;i<l;i++){
	        	if(val.indexOf(keyword[i])>-1){
	        		alert("“"+keyword[i]+"”是敏感词，请修改");
	        		$(this).prevAll(".input").click();
	        		return returnValue = false;
	        	}
	        }
		}
	});
	if(returnValue){
		$('#newsform').get(0).submit();
	}
}