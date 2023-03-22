<script>
  import {
    RequiresUpdate,
    DownloadVPN,
    DownloadStatus,
  } from "../wailsjs/go/main/App.js";

  let buttonText = "";

  let buttonAction = "";

  let vpnInstalled = false;
  // Get VPN status from Go
  RequiresUpdate().then((result) => {
    vpnInstalled = !result;
    if (!vpnInstalled) {
      buttonText = "Install";
      buttonAction = "Install";
    } else {
      buttonText = "VPN";
      buttonAction = "VPN";
    }
  });

  let vpnActive = false;

  function handleClick() {
    if (buttonAction == "VPN") {
      vpnActive = !vpnActive;
      // Change color of power button
      if (vpnActive) {
        document.querySelector(".vpn-container").style.backgroundColor =
          "#447a47";
      } else {
        document.querySelector(".vpn-container").style.backgroundColor =
          "#b44141";
      }
    } else if (buttonAction == "Install") {
      // Install VPN
      console.log("Installing VPN");
      DownloadVPN().then((result) => {
        console.log(result);
        if (result) {
          buttonText = "Downloading";
          buttonAction = "";
        }
        // Keep checking if download status
        let interval = setInterval(() => {
          DownloadStatus().then((result) => {
            if (result == 100) {
              buttonText = "VPN";
              buttonAction = "VPN";
              console.log("Done");
              clearInterval(interval);
              return;
            }
            // Round result to 2 decimal places
            result = Math.round((result + Number.EPSILON) * 100) / 100;
            buttonText = result + "%";
            console.log(result);
          });
        }, 1000);
      });
    }
  }
</script>

<main>
  <div class="main">
    <h1>VPN</h1>
    <br />
    <div class="vpn-container" on:click={handleClick} on:keypress={handleClick}>
      <div class="power-button">
        <span class="button-text">{buttonText}</span>
      </div>
    </div>
  </div>
</main>
