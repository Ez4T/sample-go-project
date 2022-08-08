import axios from "axios";

async function fetchGet(url, token) {
  const config = {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Barear ${token}`,
    },
  };
  const results = await axios.get(url, config);
  return results;
}
 
async function fetchPost(url, data, token) {
  const config = {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Barear ${token}`,
    },
  };
  const results = await axios.post(url, data, config);
  return results;
}

async function postUpload(url, data, token) {
  const config = {
    headers: {
      "Content-Type": "multipart/form-data",
      Authorization: `Barear ${token}`,
    },
  };
  const results = await axios.post(url, data, config);
  return results;
}

export { fetchGet, fetchPost, postUpload };
