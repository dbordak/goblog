function toCategoryPage(){
  var dropdown = document.getElementById('side_dd');
  if(dropdown.value != ""){
    window.location.href = "/" + dropdown.value;
  }
}
