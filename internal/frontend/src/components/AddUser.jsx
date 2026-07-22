import { useState } from "react";
import api from "../api/axios";
import Modal from "./common/Modal";

function AddUserModal({ open, onClose, onSuccess }) {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    phone: "",
    password: "",
    picture: null,
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  if (!open) return null;

  const handleChange = (e) => {
    const { name, value, files, type } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]: type === "file" ? files[0] : value,
    }));
  };

  const resetForm = () => {
    setFormData({
      username: "",
      email: "",
      phone: "",
      password: "",
      picture: null,
    });

    setError("");
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    setLoading(true);
    setError("");

    try {
      const data = new FormData();

      data.append("username", formData.username);
      data.append("email", formData.email);
      data.append("phone", formData.phone);
      data.append("password", formData.password);

      if (formData.picture) {
        data.append("picture", formData.picture);
      }

      const response = await api.post("/users", data);

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
        <div className="mb-4 rounded-lg bg-red-50 p-3 text-sm text-red-600">
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

        <div>
          <label className="block mb-2 text-sm font-medium">Picture</label>

          <input
            type="file"
            name="picture"
            accept="image/*"
            onChange={handleChange}
            className="w-full border rounded-lg px-3 py-2"
            required
          />

          {formData.picture && (
            <div className="mt-3">
              <img
                src={URL.createObjectURL(formData.picture)}
                alt="Preview"
                className="w-24 h-24 object-cover rounded-lg border"
              />
            </div>
          )}
        </div>

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
