import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/axios";
import UserTable from "../components/UserTable";
import AddUserModal from "../components/AddUser";
import EditUserModal from "../components/EditUser";
import DeleteConfirmModal from "../components/ConfirmDelete";

function UsersPage() {
  const [users, setUsers] = useState([]);
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [showAddModal, setShowAddModal] = useState(false);
  const [showEditModal, setShowEditModal] = useState(false);
  const [showDeleteModal, setShowDeleteModal] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);
  const [toast, setToast] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token");

    // Jika belum login
    if (!token) {
      navigate("/login", { replace: true });
      return;
    }

    const fetchUsers = async () => {
      try {
        const response = await api.get("/users");
        setUsers(response.data.result);
      } catch (err) {
        console.error("Error fetching users:", err);

        // Jika token tidak valid atau dihapus
        if (err.response?.status === 401) {
          localStorage.removeItem("token");
          navigate("/login", { replace: true });
          return;
        }

        setError(err.response?.data?.message || "Gagal memuat data pengguna.");
      } finally {
        setLoading(false);
      }
    };

    fetchUsers();
  }, [navigate]);

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login", { replace: true });
  };

  return (
    <div className="min-h-screen bg-slate-50 px-4 sm:px-8 py-10">
      <div className="max-w-5xl mx-auto">
        <div className="flex items-center justify-between mb-6">
          <div>
            <h1 className="text-2xl font-bold text-slate-800">
              User Management
            </h1>

            <p className="text-slate-500 text-sm mt-1">
              Kelola daftar pengguna
            </p>
          </div>

          <button
            onClick={() => setShowAddModal(true)}
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-xl cursor-pointer"
          >
            + Add User
          </button>
        </div>

        <div className="bg-white rounded-2xl shadow-xl shadow-slate-200/50 border border-slate-100 overflow-hidden mb-8">
          {loading ? (
            <div className="p-10 text-center text-slate-400">
              Memuat data...
            </div>
          ) : error ? (
            <div className="p-10 text-center text-red-500">{error}</div>
          ) : users.length === 0 ? (
            <div className="p-10 text-center text-slate-400">
              Belum ada pengguna.
            </div>
          ) : (
            <UserTable
              users={users}
              onEdit={(user) => {
                setSelectedUser(user);
                setShowEditModal(true);
              }}
              onDelete={(id) => {
                const user = users.find((u) => u.id === id);

                setSelectedUser(user);
                setShowDeleteModal(true);
              }}
            />
          )}
        </div>
        <button
          onClick={handleLogout}
          className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-xl cursor-pointer"
        >
          Logout
        </button>
      </div>

      {/* Add User Modal */}
      <AddUserModal
        open={showAddModal}
        onClose={() => setShowAddModal(false)}
        onSuccess={(newUser) => {
          setUsers((prev) => [...prev, newUser]);
        }}
      />
      {/* Edit User Modal */}
      <EditUserModal
        open={showEditModal}
        user={selectedUser}
        onClose={() => {
          setShowEditModal(false);
          setSelectedUser(null);
        }}
        onSuccess={(updatedUser) => {
          setUsers((prev) =>
            prev.map((user) =>
              user.id === updatedUser.id ? updatedUser : user,
            ),
          );
        }}
      />
      {/* Delete Confirm Modal */}
      <DeleteConfirmModal
        open={showDeleteModal}
        user={selectedUser}
        onClose={() => {
          setShowDeleteModal(false);
          setSelectedUser(null);
        }}
        onSuccess={(id) => {
          setUsers((prev) => prev.filter((user) => user.id !== id));
        }}
      />

      {/* Toast notification */}
      {toast && (
        <div
          className={`fixed bottom-6 right-6 z-50 flex items-center gap-3 px-4 py-3 rounded-xl shadow-lg text-sm font-medium text-white transition-all ${
            toast.type === "success" ? "bg-emerald-600" : "bg-red-600"
          }`}
        >
          <span>{toast.type === "success" ? "✓" : "✕"}</span>
          <span>{toast.message}</span>
          <button
            onClick={() => setToast(null)}
            className="ml-2 text-white/80 hover:text-white cursor-pointer"
          >
            ✕
          </button>
        </div>
      )}
    </div>
  );
}

export default UsersPage;
