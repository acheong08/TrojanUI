<script>
  import { RequiresUpdate } from "../wailsjs/go/main/App.js";

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
