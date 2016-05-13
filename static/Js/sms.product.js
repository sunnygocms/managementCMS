$(document).ready(function() {
	$("#newsform").validate({
		ignore : "",
		rules : {
			name : "required",
			description : "required",
			sample : "required",
			tvod_price:{required:true,number:true},
			svod_price:{required:true,number:true},
			tvod_sort:{required:true,number:true,range:[1,99]},
			svod_sort:{required:true,number:true,range:[1,99]}

		},
		messages : {
			name : "请输入分类标题",
			description : "请输入描述",
			sample : "请输入样例",
			tvod_price:{
				required:"输入不能为空",
				number:"只能够输入数值"
			},
			svod_price:{
				required:"输入不能为空",
				number:"只能够输入数值"
			},
			tvod_sort:{
				required:"输入不能为空",
				number:"只能够输入数值",
				range:"取值范围在1-99"
			},
			svod_sor:{
					required:"输入不能为空",
					number:"只能够输入数值",
					range:"取值范围在1-99"
			}
		}
	});

	$('#newsform').submit(function() {
		if ($("#newsform").valid()) {
			return true;
		}

	});
	
	$("#svod").change(function(){
		 if(!$(this).prop("checked")){
			 $("#svod_price").prop("disabled","disabled"); 
		 }else{
			 $("#svod_price").prop("disabled",""); 
		 }
	});
	$("#tvod").change(function(){
		 if(!$(this).prop("checked")){
			 $("#tvod_price").prop("disabled","disabled"); 
		 }else{
			 $("#tvod_price").prop("disabled",""); 
		 }
	});
	
	$("#channel_id").change(function(){
			$('#searchForm').submit();
	});
	
});