$(document).ready(function() {
	$("#recommendListForm").submit(function(){
		//return $(this).find("input[name='checkid[]']:checked").length?true:false;
		return $("form:last").find("input[name='checkid[]']:checked").length?true:false;
	});
});

function checkall(){
	if($("#CheckAll").prop("checked")){
		$("#recommendListForm input[name='checkid[]']").prop("checked",true);
	}else{
		$("#recommendListForm input[name='checkid[]']").prop("checked",false);
	}
	
}