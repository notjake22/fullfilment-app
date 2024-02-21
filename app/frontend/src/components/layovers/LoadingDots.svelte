<div class="dots-9"></div>

<style>
    .dots-9 {
        width: 15px;
        aspect-ratio: 1;
        border-radius: 50%;
        clip-path: inset(-45px);
        box-shadow: -60px 15px,-60px 15px,-60px 15px;
        transform: translateY(-15px);
        animation: d9 1s infinite linear;
    }

    @keyframes d9{
        16.67% {box-shadow:-60px 15px,-60px 15px,19px 15px}
        33.33% {box-shadow:-60px 15px,  0px 15px,19px 15px}
        40%,60%{box-shadow:-19px 15px,  0px 15px,19px 15px}
        66.67% {box-shadow:-19px 15px,  0px 15px,60px 15px}
        83.33% {box-shadow:-19px 15px, 60px 15px,60px 15px}
        100%   {box-shadow: 60px 15px, 60px 15px,60px 15px}
    }
</style>