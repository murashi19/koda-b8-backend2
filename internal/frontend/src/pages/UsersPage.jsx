import { useEffect, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import api from "../api/axios";
import UserTable from "../components/UserTable";
import AddUserModal from "../components/AddUser";
import EditUserModal from "../components/EditUser";
import DeleteConfirmModal from "../components/ConfirmDelete";
import UploadPicture from "../components/UploadPicture";
import Search from "../components/Search";

function UsersPage() {
  const [users, setUsers] = useState([]);
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  // Modal
  const [showAddModal, setShowAddModal] = useState(false);
  const [showEditModal, setShowEditModal] = useState(false);
  const [showDeleteModal, setShowDeleteModal] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);
  const [showUploadModal, setShowUploadModal] = useState(false);
  // Notif
  const [toast, setToast] = useState(null);

  // Params
  const [searchParams, setSearchParams] = useSearchParams();
  const keyword = searchParams.get("q") || "";

  // Pagination — single source of truth: page/limit come from the URL,
  // total/totalPage come from the API response. No more separate
  // "pageInfo" state that never got synced.
  const page = Number(searchParams.get("page")) || 1;
  const limit = Number(searchParams.get("limit")) || 5;
  const [total, setTotal] = useState(0);
  const [totalPage, setTotalPage] = useState(1);

  useEffect(() => {
    const accessToken = localStorage.getItem("access_token");

    if (!accessToken) {
      navigate("/login", { replace: true });
      return;
    }

    // Fetch Data
    const fetchUsers = async () => {
      try {
        setLoading(true);

        const params = { page, limit };

        if (keyword.trim()) {
          params[`search[username]`] = keyword.trim();
        }

        const response = await api.get("/users", {
          params,
        });

        setUsers(response.data.result.data);
        setTotal(response.data.result.total);
        setTotalPage(response.data.result.total_page);
        setError("");
      } catch (err) {
        console.error(err);

        setError(err.response?.data?.message || "Gagal memuat data pengguna.");
      } finally {
        setLoading(false);
      }
    };
    fetchUsers();
  }, [page, limit, keyword, navigate]);

  // Search Input
  const handleInput = (e) => {
    const value = e.target.value;
    const params = { page: 1, limit };
    if (value) params.q = value;
    setSearchParams(params);
  };

  const handleSearch = (e) => {
    e.preventDefault();
  };

  // Logout
  const handleLogout = () => {
    localStorage.removeItem("access_token");
    localStorage.removeItem("refresh_token");
    navigate("/login", { replace: true });
  };

  // Button Pagination
  const handlePrev = () => {
    if (page <= 1) return;

    setSearchParams({
      page: page - 1,
      limit,
      ...(keyword && { q: keyword }),
    });
  };

  const handleNext = () => {
    if (page >= totalPage) return;

    setSearchParams({
      page: page + 1,
      limit,
      ...(keyword && { q: keyword }),
    });
  };

  return (
    <div className="min-h-screen bg-slate-50 px-4 sm:px-8 py-10">
      <div className="max-w-5xl mx-auto">
        <div className="flex items-center justify-between mb-6 gap-10">
          <div>
            <h1 className="text-2xl font-bold text-slate-800">
              User Management
            </h1>

            <p className="text-slate-500 text-sm mt-1">
              Kelola daftar pengguna
            </p>
          </div>
          <div className="w-1/2">
            <Search
              keyword={keyword}
              onChange={handleInput}
              onSubmit={handleSearch}
            />
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
              onUpload={(user) => {
                setSelectedUser(user);
                setShowUploadModal(true);
              }}
            />
          )}

          {/* Pagination footer — attached to the same card as the table */}
          {!loading && !error && users.length > 0 && (
            <div className="flex items-center justify-between gap-4 px-6 py-4 border-t border-slate-100 bg-slate-50/50">
              <span className="text-sm text-slate-500">
                Halaman{" "}
                <span className="font-medium text-slate-700">{page}</span> dari{" "}
                <span className="font-medium text-slate-700">{totalPage}</span>{" "}
                <span className="text-slate-400">({total} total)</span>
              </span>

              <div className="flex items-center gap-2">
                <button
                  disabled={page === 1}
                  onClick={handlePrev}
                  className="px-3 py-1.5 text-sm border border-slate-200 bg-white rounded-lg hover:bg-slate-50 disabled:opacity-50 disabled:hover:bg-white cursor-pointer disabled:cursor-not-allowed"
                >
                  Previous
                </button>

                <button
                  disabled={page >= totalPage}
                  onClick={handleNext}
                  className="px-3 py-1.5 text-sm border border-slate-200 bg-white rounded-lg hover:bg-slate-50 disabled:opacity-50 disabled:hover:bg-white cursor-pointer disabled:cursor-not-allowed"
                >
                  Next
                </button>
              </div>
            </div>
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

      <UploadPicture
        open={showUploadModal}
        user={selectedUser}
        onClose={() => {
          setShowUploadModal(false);
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
