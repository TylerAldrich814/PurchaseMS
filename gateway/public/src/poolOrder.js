const urlParams = new URLSearchParams(window.location.search);
const customerID = urlParams.get("customerID");
const orderID = urlParams.get("orderID");

const order = {
  customerID,
  orderID,
  status: "pending",
};

async function poolOrderStatus() {
  console.log(`poolORderStatus: ${order.status}`)
  const response = await fetch(`api/customers/${customerID}/orders/${orderID}`)

  const data = await response.json();
  console.log(data)
  console.log(`Link: "${data.PaymentLink}"`)

  if( data.status === "waiting_payment" ){
    order.status = "Your Order is waiting for payment...";
    document.getElementById("orderStatus").innerText = order.status;

    document.querySelector(".payment-popup").style.display = 'block';
    document.getElementById("payment-link").href = data.paymentLink
  }

  if( data.Status === "paid" ){
    order.status = "Your order has been paid for! Please wait while its being processed.."
    docuemnt.getElementById("orderStatus").innerText = order.status;

    setTimeout(poolOrderStatus, 5000);
  } else if( data.status === "ready" ){
    order.status = "Ready";
    document.querySelector(".payment-popup").style.display = "none";
    
    document.querySelector(".ready-popup").style.display = "block";
    document.getElementById("orderID").innerText = orderID;
    document.getElementById("orderStatus").innerText = order.status
  } else {
    setTimeout(poolOrderStatus, 5000);
  }
}

poolOrderStatus();
