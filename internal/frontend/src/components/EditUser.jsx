import { useState } from "react";
import api from "../api/axios";
import Modal from "./common/Modal";

function EditUserModal({ open, user, onClose, onSuccess }) {
  const [formData, setFormData] = useState(() => ({
    username: user?.username || "",
    email: user?.email || "",
    phone: user?.phone || "",
  }));

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleChange = (e) => {
    const { name, value } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    setLoading(true);
    setError("");

    try {
      const response = await api.patch(`/users/${user.id}`, formData);

      onSuccess(response.data.result);

      onClose();
    } catch (err) {
      setError(err.response?.data?.message || "Gagal mengupdate user.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal open={open} title="Edit User" onClose={onClose}>
      {error && (
        <div className="mb-4 rounded-lg bg-red-100 text-red-600 p-3">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-4">
        <label htmlFor="username">Username</label>
        <input
          type="text"
          name="username"
          value={formData.username}
          onChange={handleChange}
          className="w-full border rounded-lg px-3 py-2"
          required
        />

        <label htmlFor="email">Email</label>
        <input
          type="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
          className="w-full border rounded-lg px-3 py-2"
          required
        />

        <label htmlFor="phone">Phone</label>
        <input
          type="text"
          name="phone"
          value={formData.phone}
          onChange={handleChange}
          className="w-full border rounded-lg px-3 py-2"
          required
        />

        <div className="flex justify-end gap-3">
          <button
            type="button"
            onClick={onClose}
            className="border px-4 py-2 rounded-lg"
          >
            Cancel
          </button>

          <button
            type="submit"
            disabled={loading}
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg"
          >
            {loading ? "Updating..." : "Update"}
          </button>
        </div>
      </form>
    </Modal>
  );
}

export default EditUserModal;
