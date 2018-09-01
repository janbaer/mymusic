function buildMP3DBApiUrl() {
  if (window.location.hostname !== 'localhost') {
    return `${window.location.protocol}//${window.location.hostname}:${window.location.port}/api`;
  }
  return `${window.location.protocol}//${window.location.hostname}:8082`;
}

export {
  buildMP3DBApiUrl
};
