import { useState } from "react";
import api from "../api/axios";
import Modal from "./common/Modal";

function AddUserModal({ open, onClose, onSuccess }) {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    phone: "",
    password: "",
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  if (!open) return null;

  const handleChange = (e) => {
    const { name, value } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const resetForm = () => {
    setFormData({
      username: "",
      email: "",
      phone: "",
      password: "",
    });

    setError("");
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    setLoading(true);
    setError("");

    try {
      const response = await api.post("/auth/register", formData);

      onSuccess(response.data.result);

      resetForm();
      onClose();
    } catch (err) {
      setError(err.response?.data?.message || "Gagal menambahkan user.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal
      open={open}
      title="Add User"
      onClose={() => {
        resetForm();
        onClose();
      }}
    >
      {error && (
        <div className="mb-4 rounded-lg bg-red-50 text-red-600 p-3 text-sm">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-4">
        <input
          type="text"
          name="username"
          placeholder="Username"
          value={formData.username}
          onChange={handleChange}
          className="w-full border rounded-lg px-3 py-2"
          required
        />

        <input
          type="email"
          name="email"
          placeholder="Email"
          value={formData.email}
          onChange={handleChange}
          className="w-full border rounded-lg px-3 py-2"
          required
        />

        <input
          type="text"
          name="phone"
          placeholder="Phone"
          value={formData.phone}
          onChange={handleChange}
          className="w-full border rounded-lg px-3 py-2"
          required
        />

        <input
          type="password"
          name="password"
          placeholder="Password"
          value={formData.password}
          onChange={handleChange}
          className="w-full border rounded-lg px-3 py-2"
          required
        />

        <div className="flex justify-end gap-3 pt-3">
          <button
            type="button"
            onClick={() => {
              resetForm();
              onClose();
            }}
            className="px-4 py-2 border rounded-lg cursor-pointer"
          >
            Cancel
          </button>

          <button
            type="submit"
            disabled={loading}
            className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg disabled:opacity-50 cursor-pointer"
          >
            {loading ? "Saving..." : "Save"}
          </button>
        </div>
      </form>
    </Modal>
  );
}

export default AddUserModal;
