import { useState, useEffect } from "react";
import api from "../api/axios";
import Modal from "./common/Modal";

function EditUserModal({ open, user, onClose, onSuccess }) {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    phone: "",
  });

  const [fetching, setFetching] = useState(false);
  const [updating, setUpdating] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!open || !user?.id) return;

    let ignore = false;

    const fetchUser = async () => {
      try {
        setFetching(true);

        const response = await api.get(`/users/${user.id}`);

        if (ignore) return;

        const data = response.data.result;

        setFormData({
          username: data.username ?? "",
          email: data.email ?? "",
          phone: data.phone ?? "",
        });

        setError("");
      } catch (err) {
        if (ignore) return;

        console.error(err);

        setError(err.response?.data?.message ?? "Gagal mengambil data user.");
      } finally {
        if (!ignore) {
          setFetching(false);
        }
      }
    };

    fetchUser();

    return () => {
      ignore = true;
    };
  }, [open, user?.id]);

  const handleChange = (e) => {
    const { name, value } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleClose = () => {
    setFormData({
      username: "",
      email: "",
      phone: "",
    });

    setError("");

    onClose();
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!user?.id) return;

    try {
      setUpdating(true);
      setError("");

      const response = await api.patch(`/users/${user.id}`, formData);

      onSuccess(response.data.result);

      handleClose();
    } catch (err) {
      console.error(err);

      setError(err.response?.data?.message ?? "Gagal mengupdate user.");
    } finally {
      setUpdating(false);
    }
  };
  return (
    <Modal open={open} title="Edit User" onClose={handleClose}>
      {error && (
        <div className="mb-4 rounded-lg bg-red-100 p-3 text-red-600">
          {error}
        </div>
      )}

      {fetching ? (
        <div className="py-8 text-center">Loading user...</div>
      ) : (
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="mb-1 block">Username</label>

            <input
              type="text"
              name="username"
              value={formData.username}
              onChange={handleChange}
              disabled={updating}
              className="w-full rounded-lg border px-3 py-2"
              required
            />
          </div>

          <div>
            <label className="mb-1 block">Email</label>

            <input
              type="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
              disabled={updating}
              className="w-full rounded-lg border px-3 py-2"
              required
            />
          </div>

          <div>
            <label className="mb-1 block">Phone</label>

            <input
              type="text"
              name="phone"
              value={formData.phone}
              onChange={handleChange}
              disabled={updating}
              className="w-full rounded-lg border px-3 py-2"
              required
            />
          </div>

          <div className="flex justify-end gap-3">
            <button
              type="button"
              onClick={handleClose}
              disabled={updating}
              className="rounded-lg border px-4 py-2"
            >
              Cancel
            </button>

            <button
              type="submit"
              disabled={updating}
              className="rounded-lg bg-blue-600 px-4 py-2 text-white hover:bg-blue-700 disabled:cursor-not-allowed disabled:opacity-50"
            >
              {updating ? "Updating..." : "Update"}
            </button>
          </div>
        </form>
      )}
    </Modal>
  );
}

export default EditUserModal;
