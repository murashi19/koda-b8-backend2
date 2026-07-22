import { useState } from "react";
import api from "../api/axios";
import Modal from "./common/Modal";

function UploadPicture({ open, user, onClose, onSuccess }) {
  const [file, setFile] = useState(null);
  const [loading, setLoading] = useState(false);

  if (!user) return null;

  const handleUpload = async () => {
    if (!file) {
      alert("Please select a file to upload.");
      return;
    }

    setLoading(true);

    try {
      const formData = new FormData();
      formData.append("picture", file);

      const response = await api.patch(`/users/${user.id}/upload`, formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });

      console.log(response.data.result);
      onSuccess(response.data.result);
      onClose();
    } catch (err) {
      alert(err.response?.data?.message ?? "Failed to upload picture.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal
      open={open}
      title="Upload Picture"
      onClose={onClose}
      width="max-w-sm"
    >
      <div className="space-y-5">
        <input
          type="file"
          accept="image/*"
          onChange={(e) => setFile(e.target.files[0])}
        />

        <div className="flex justify-end gap-3">
          <button onClick={onClose} className="border px-4 py-2 rounded-lg">
            Cancel
          </button>

          <button
            onClick={handleUpload}
            disabled={loading}
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg"
          >
            {loading ? "Uploading..." : "Upload"}
          </button>
        </div>
      </div>
    </Modal>
  );
}

export default UploadPicture;
