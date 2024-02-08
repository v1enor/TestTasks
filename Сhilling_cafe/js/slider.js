$(document).ready(function(){
    $(".carousel-container").slick({
        autoplay: true,
        slidesToShow: 3,
        slidesToScroll: 1,
        speed: 350,
        responsive: [
            {
                breakpoint: 550, 
                settings: {
                    slidesToShow: 1, 
                    slidesToScroll: 1
                }
            }
        ]
    });
});