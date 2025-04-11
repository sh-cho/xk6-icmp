import icmp from "k6/x/icmp";

export default function () {
  let p = icmp.ping("www.google.com");
  console.log(p);
}
