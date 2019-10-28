import * as React from 'react';

import '../style/Loading';
class Loading extends React.Component {
    render() {
        return(
            <div id='loading'>
                <svg width="611" height="133" viewBox="0 0 611 133" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M18.7617 78.668V128H3.13428V5.00781H43.0054C48.5243 5.12044 53.7334 5.99333 58.6328 7.62646C63.5885 9.2596 67.9248 11.6248 71.6416 14.7222C75.3584 17.8195 78.2868 21.6489 80.4268 26.2104C82.623 30.772 83.7212 36.0093 83.7212 41.9224C83.7212 47.8354 82.623 53.0728 80.4268 57.6343C78.2868 62.1395 75.3584 65.9408 71.6416 69.0381C67.9248 72.1354 63.5885 74.5007 58.6328 76.1338C53.7334 77.7669 48.5243 78.6117 43.0054 78.668H18.7617ZM18.7617 65.8281H43.0054C46.6095 65.7718 49.9321 65.2087 52.9731 64.1387C56.0142 63.0124 58.661 61.4355 60.9136 59.4082C63.1662 57.3809 64.9119 54.9312 66.1509 52.0591C67.4461 49.1307 68.0938 45.8081 68.0938 42.0913C68.0938 38.3745 67.4461 35.0238 66.1509 32.0391C64.9119 29.0544 63.1662 26.5202 60.9136 24.4365C58.7173 22.3529 56.0705 20.7479 52.9731 19.6216C49.9321 18.4953 46.6095 17.904 43.0054 17.8477H18.7617V65.8281Z" stroke="#2CB256" stroke-width="6"/>
                    <path d="M146.315 77.8232H121.734V128H106.191V5.00781H142.261C148.005 5.12044 153.411 5.93701 158.479 7.45752C163.548 8.97803 167.969 11.2588 171.742 14.2998C175.571 17.3408 178.556 21.1702 180.696 25.7881C182.892 30.3496 183.99 35.7277 183.99 41.9224C183.99 45.9207 183.399 49.5812 182.216 52.9038C181.09 56.2264 179.513 59.2393 177.486 61.9424C175.458 64.6455 173.037 67.0389 170.221 69.1226C167.405 71.2062 164.308 72.9801 160.929 74.4443L187.031 126.986L186.947 128H170.475L146.315 77.8232ZM121.734 64.9834H142.683C146.175 64.9271 149.469 64.3921 152.566 63.3784C155.664 62.3084 158.367 60.7879 160.676 58.8169C163.041 56.8459 164.899 54.4525 166.251 51.6367C167.603 48.7646 168.278 45.4702 168.278 41.7534C168.278 37.8114 167.631 34.3761 166.335 31.4478C165.04 28.4631 163.238 25.9852 160.929 24.0142C158.62 21.9868 155.861 20.4663 152.651 19.4526C149.497 18.439 146.034 17.904 142.261 17.8477H121.734V64.9834Z" stroke="#2CB256" stroke-width="6"/>
                    <path d="M277.164 71.1499H225.804V114.738H285.695V128H210.177V5.00781H284.935V18.3545H225.804V57.8877H277.164V71.1499Z" stroke="#2CB256" stroke-width="6"/>
                    <path d="M377.771 96.9141C377.771 93.141 376.898 89.9591 375.152 87.3687C373.462 84.7782 371.266 82.61 368.563 80.8643C365.86 79.0622 362.875 77.5698 359.609 76.3872C356.399 75.2046 353.302 74.1346 350.317 73.1772C345.981 71.7694 341.56 70.0799 337.055 68.1089C332.606 66.1379 328.523 63.7163 324.806 60.8442C321.146 57.9722 318.133 54.5933 315.768 50.7075C313.459 46.7655 312.304 42.1476 312.304 36.854C312.304 31.5604 313.459 26.8299 315.768 22.6626C318.133 18.4953 321.202 14.9756 324.975 12.1035C328.748 9.23145 333.028 7.06331 337.815 5.59912C342.602 4.07861 347.417 3.31836 352.26 3.31836C357.61 3.31836 362.763 4.2194 367.718 6.02148C372.674 7.76725 377.067 10.2451 380.896 13.4551C384.725 16.665 387.795 20.5508 390.104 25.1123C392.412 29.6738 393.623 34.7703 393.736 40.4019H377.686C377.236 36.854 376.363 33.644 375.067 30.772C373.772 27.8436 372.055 25.3376 369.915 23.2539C367.775 21.1702 365.212 19.5653 362.228 18.439C359.299 17.2563 355.977 16.665 352.26 16.665C349.275 16.665 346.347 17.0874 343.475 17.9321C340.659 18.7769 338.125 20.0439 335.872 21.7334C333.676 23.4229 331.902 25.5065 330.55 27.9844C329.255 30.4622 328.607 33.3343 328.607 36.6006C328.664 40.1484 329.565 43.1613 331.311 45.6392C333.056 48.0607 335.253 50.1162 337.899 51.8057C340.603 53.4951 343.503 54.903 346.6 56.0293C349.754 57.1556 352.71 58.1411 355.47 58.9858C358.511 59.9432 361.58 61.0413 364.677 62.2803C367.775 63.4629 370.759 64.8426 373.631 66.4194C376.503 67.9963 379.178 69.7702 381.656 71.7412C384.134 73.7122 386.302 75.9648 388.161 78.499C390.019 80.9769 391.455 83.7363 392.469 86.7773C393.539 89.762 394.074 93.0846 394.074 96.7451C394.074 102.264 392.835 107.079 390.357 111.19C387.935 115.301 384.782 118.736 380.896 121.496C377.01 124.199 372.618 126.254 367.718 127.662C362.819 129.014 357.919 129.689 353.02 129.689C347.557 129.689 342.208 128.845 336.97 127.155C331.733 125.466 327.031 123.044 322.863 119.891C318.752 116.681 315.402 112.795 312.811 108.233C310.221 103.616 308.841 98.3783 308.672 92.5215H324.637C325.144 96.3509 326.186 99.758 327.763 102.743C329.34 105.671 331.367 108.177 333.845 110.261C336.323 112.288 339.167 113.837 342.376 114.907C345.643 115.92 349.191 116.427 353.02 116.427C356.061 116.427 359.046 116.061 361.974 115.329C364.959 114.541 367.606 113.358 369.915 111.781C372.223 110.148 374.11 108.121 375.574 105.699C377.038 103.221 377.771 100.293 377.771 96.9141Z" stroke="#2CB256" stroke-width="6"/>
                    <path d="M481.672 96.9141C481.672 93.141 480.799 89.9591 479.053 87.3687C477.364 84.7782 475.167 82.61 472.464 80.8643C469.761 79.0622 466.777 77.5698 463.51 76.3872C460.3 75.2046 457.203 74.1346 454.218 73.1772C449.882 71.7694 445.461 70.0799 440.956 68.1089C436.507 66.1379 432.424 63.7163 428.708 60.8442C425.047 57.9722 422.034 54.5933 419.669 50.7075C417.36 46.7655 416.206 42.1476 416.206 36.854C416.206 31.5604 417.36 26.8299 419.669 22.6626C422.034 18.4953 425.103 14.9756 428.876 12.1035C432.65 9.23145 436.93 7.06331 441.716 5.59912C446.503 4.07861 451.318 3.31836 456.161 3.31836C461.511 3.31836 466.664 4.2194 471.62 6.02148C476.575 7.76725 480.968 10.2451 484.797 13.4551C488.627 16.665 491.696 20.5508 494.005 25.1123C496.314 29.6738 497.525 34.7703 497.637 40.4019H481.587C481.137 36.854 480.264 33.644 478.969 30.772C477.674 27.8436 475.956 25.3376 473.816 23.2539C471.676 21.1702 469.114 19.5653 466.129 18.439C463.201 17.2563 459.878 16.665 456.161 16.665C453.176 16.665 450.248 17.0874 447.376 17.9321C444.56 18.7769 442.026 20.0439 439.773 21.7334C437.577 23.4229 435.803 25.5065 434.452 27.9844C433.156 30.4622 432.509 33.3343 432.509 36.6006C432.565 40.1484 433.466 43.1613 435.212 45.6392C436.958 48.0607 439.154 50.1162 441.801 51.8057C444.504 53.4951 447.404 54.903 450.501 56.0293C453.655 57.1556 456.612 58.1411 459.371 58.9858C462.412 59.9432 465.481 61.0413 468.579 62.2803C471.676 63.4629 474.661 64.8426 477.533 66.4194C480.405 67.9963 483.08 69.7702 485.558 71.7412C488.035 73.7122 490.204 75.9648 492.062 78.499C493.92 80.9769 495.356 83.7363 496.37 86.7773C497.44 89.762 497.975 93.0846 497.975 96.7451C497.975 102.264 496.736 107.079 494.258 111.19C491.837 115.301 488.683 118.736 484.797 121.496C480.912 124.199 476.519 126.254 471.62 127.662C466.72 129.014 461.821 129.689 456.921 129.689C451.459 129.689 446.109 128.845 440.872 127.155C435.634 125.466 430.932 123.044 426.765 119.891C422.654 116.681 419.303 112.795 416.712 108.233C414.122 103.616 412.742 98.3783 412.573 92.5215H428.539C429.045 96.3509 430.087 99.758 431.664 102.743C433.241 105.671 435.268 108.177 437.746 110.261C440.224 112.288 443.068 113.837 446.278 114.907C449.544 115.92 453.092 116.427 456.921 116.427C459.962 116.427 462.947 116.061 465.875 115.329C468.86 114.541 471.507 113.358 473.816 111.781C476.125 110.148 478.011 108.121 479.476 105.699C480.94 103.221 481.672 100.293 481.672 96.9141Z" stroke="#2CB256" stroke-width="6"/>
                    <path d="M559.302 53.2417L585.827 5.00781H604.242L568.51 65.9971L605.086 128H586.84L559.64 78.8369L532.355 128H513.856L532.187 96.9985L550.517 65.9971L514.785 5.00781H533.116L559.302 53.2417Z" stroke="#2CB256" stroke-width="6"/>
                </svg>

                <p>Loading...</p>    
            </div>
        );
    
    }  
}  
export default Loading;

