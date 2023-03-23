<script>
  import {
    RequiresUpdate,
    DownloadVPN,
    DownloadStatus,
    StartVPN,
    StopVPN,
  } from "../wailsjs/go/main/App.js";

  let buttonText = "Loading";
  let buttonAction = "";
  let vpnInstalled = false;
  let vpnActive = false;
  let vpnContainerBgColor = "#BF4630";

  async function initialize() {
    const result = await RequiresUpdate();
    vpnInstalled = !result;
    buttonText = vpnInstalled ? "VPN" : "Install";
    buttonAction = vpnInstalled ? "VPN" : "Install";
  }

  initialize();

  async function handleClick() {
    if (buttonAction === "VPN") {
      let success;
      if (vpnActive) {
        success = await StopVPN();
      } else {
        success = await StartVPN();
      }
      if (!success) {
        buttonText = "Error";
        return;
      }
      buttonText = "VPN";
      vpnActive = !vpnActive;
      vpnContainerBgColor = vpnActive ? "#238C47" : "#BF4630";
    } else if (buttonAction === "Install") {
      console.log("Installing VPN");
      const result = await DownloadVPN();

      if (result) {
        buttonText = "Downloading";
        buttonAction = "";

        let interval = setInterval(async () => {
          const status = await DownloadStatus();
          if (status === 100) {
            buttonText = "VPN";
            buttonAction = "VPN";
            console.log("Done");
            clearInterval(interval);
            return;
          }
          const roundedStatus =
            Math.round((status + Number.EPSILON) * 100) / 100;
          buttonText = `${roundedStatus}%`;
          console.log(roundedStatus);
        }, 1000);
      }
    }
  }
</script>

<main>
  <div class="main">
    <h1>Boredom</h1>
    <br />
    <div
      class="vpn-container"
      on:click={handleClick}
      on:keypress={handleClick}
      style="background-color: {vpnContainerBgColor};"
    >
      <div class="power-button">
        <span class="button-text">{buttonText}</span>
      </div>
    </div>
  </div>
</main>
