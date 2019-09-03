export default {
  name: "AccountDetails",
  props: {
    accountid: String,
    owner: String,
  },
  template: '<div><h2>{{ accountid }}: {{ owner }}</h2><h3>Bookings</h3></div>'
}