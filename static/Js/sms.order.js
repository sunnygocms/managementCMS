$(function(){    
    $(".grayTips").each(function(){ //遍历每个文本框
        var objTextBox=$(this);
        var oldText=$.trim(objTextBox.val());
        objTextBox.css("color","#888"); 
        objTextBox.focus(function(){
            if(objTextBox.val()!=oldText){
                objTextBox.css("color","#000");
            }
            else{
                objTextBox.val("").css("color","#888");
            }
        });
        objTextBox.blur(function(){
            if(objTextBox.val()==""){
                objTextBox.val(oldText).css("color","#888");
            }
        });
        objTextBox.keydown(function(){
            objTextBox.css("color","#000");
        });
    });
});