function SelectAll(selectAllID, checkboxClass) {
    // Check of Uncheck All checkboxes
    var selectAll = document.getElementById(selectAllID);
    var checkboxes = document.getElementsByClassName(checkboxClass);
    if(selectAll.checked){
        // Select all
        for (x = 0; x < checkboxes.length; x++) {
            checkboxes.item(x).checked = true;
        }
    } else {
        // Deselect All
        for (x = 0; x < checkboxes.length; x++) {
            checkboxes.item(x).checked = false;
        }
    }
};

$(document).ready(function(){
    // Changing state of SelectAll checkbox based on values
    $(".namedLocationCheckbox").click(function(){
        var checkboxes = document.getElementsByClassName("namedLocationCheckbox");
        if(checkboxes.length == $(".namedLocationCheckbox:checked").length) {
            $("#selectAllNamed").prop("checked", true);
        } else {
            $("#selectAllNamed").removeAttr("checked");
        }
    });

    $(".unnamedLocationCheckbox").click(function(){
        var checkboxes = document.getElementsByClassName("unnamedLocationCheckbox");
        if(checkboxes.length == $(".unnamedLocationCheckbox:checked").length) {
            $("#selectAllUnnamed").prop("checked", true);
        } else {
            $("#selectAllUnnamed").removeAttr("checked");
        }
    });
});
