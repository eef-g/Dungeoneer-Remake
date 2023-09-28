//events.js
import "../wailsjs/runtime/runtime.js"

function subscribe(eventName, listener) {
  window.runtime.EventsOn(eventName, listener);
}

function unsubscribe(eventName, listener) {
  document.removeEventListener(eventName, listener);
}

function publish(eventName, data) {
  const event = new CustomEvent(eventName, { detail: data });
  document.dispatchEvent(event);
}

export { publish, subscribe, unsubscribe};
