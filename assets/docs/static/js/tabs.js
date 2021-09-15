$('.panes').find('.pane').each(function(idx, item) {
    var tabs = $(this).closest('.code-tabs').find('.tabs'),
        title = $(this).attr('title');
    tabs.append('<li><a href="#">'+title+'</a></li');
});

$('.code-tabs ul.tabs').each(function() {
    $(this).find("li:first").addClass('active');
})

$('.code-tabs .panes').each(function() {
    $(this).find('.pane').hide()
    $(this).find("div:first").addClass('active').show();
});

$('.tabs a').click(function(e){
    e.preventDefault();
    var tab = $(this).parent(),
        tabIndex = tab.index(),
        tabPanel = $(this).closest('.code-tabs'),
        tabPane = tabPanel.find('.pane').eq(tabIndex);
    tabPanel.find('.pane.active').removeClass('active').hide();
    tabPanel.find('.active').removeClass('active');
    tab.addClass('active');
    tabPane.addClass('active').fadeIn('slow');
});