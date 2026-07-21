import { useState } from "react";
import api from "../api/axios";
import Modal from "./common/Modal";

function DeleteConfirmModal({ open, user, onClose, onSuccess }) {
  const [loading, setLoading] = useState(false);

  if (!user) return null;

  const handleDelete = async () => {
    setLoading(true);

    try {
      await api.delete(`/users/${user.id}`);

      onSuccess(user.id);

      onClose();
    } catch (err) {
      alert(err.response?.data?.message ?? "Gagal menghapus user.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal open={open} title="Delete User" onClose={onClose} width="max-w-sm">
      <div className="space-y-5">
        <p className="text-gray-600">
          Apakah kamu yakin ingin menghapus user
          <span className="font-semibold"> {user.username}</span>?
        </p>

        <div className="flex justify-end gap-3">
          <button onClick={onClose} className="border px-4 py-2 rounded-lg">
            Cancel
          </button>

          <button
            onClick={handleDelete}
            disabled={loading}
            className="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-lg"
          >
            {loading ? "Deleting..." : "Delete"}
          </button>
        </div>
      </div>
    </Modal>
  );
}

export default DeleteConfirmModal;
