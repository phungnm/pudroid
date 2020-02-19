var last_url="";
var last_custom="";
var total_url=0;
var regex = RegExp(/(https?:\/\/)?([\w\-])+\.{1}([a-zA-Z]{2,63})([\/\w-]*)*\/?\??([^#\n\r]*)?#?([^\n\r]*)/);
$(document).ready(function($) {
	
	$("#shorten-list").delegate('.shorten-copy', 'click', function(event) {
		let button =$(this);
		if(button.prop("disabled"));
		else
		{
					let ele =	button.attr("attr-target");

					var $temp = $("<input>");
$("body").append($temp);
$temp.val($("#"+ele).val()).select();
document.execCommand("copy");
$temp.remove();
	
button.prop("disabled",true);
button.html(`<i class="far fa-clipboard mr-3"></i> Copied`);
setTimeout(function(){
 button.html(`<i class="far fa-clipboard mr-3"></i> Copy `);
button.prop("disabled",false);

}, 1500);
		}


	});
	$("#shorten-submit").on('click',function(event){



	});
	$("#url-input").keyup(function(){
		let check = regex.test($(this).val());
		console.log($(this).val(),check);
		if(check)
		 $("#shorten-submit").prop("disabled",false);
		else $("#shorten-submit").prop("disabled",true);
	});
	$("#custom-input").keyup(function(){
		let check = regex.test(	$("#url-input").val());
		if(check)
		{
			$("#shorten-submit").prop("disabled",false);
		}
	});

});