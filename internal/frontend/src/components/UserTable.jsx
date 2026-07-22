function UserTable({ users, onEdit, onDelete, onUpload }) {
  if (users.length === 0) {
    return (
      <div className="p-10 text-center text-slate-400 text-sm">
        Belum ada pengguna.
      </div>
    );
  }

  return (
    <div className="overflow-x-auto">
      <table className="w-full text-left text-sm">
        <thead>
          <tr className="bg-slate-50 text-slate-500 uppercase text-xs tracking-wide">
            <th className="px-5 py-3">No</th>
            <th className="px-5 py-3">Picture</th>
            <th className="px-5 py-3">Username</th>
            <th className="px-5 py-3">Email</th>
            <th className="px-5 py-3">Phone</th>
            <th className="px-5 py-3 text-right">Actions</th>
          </tr>
        </thead>

        <tbody className="divide-y divide-slate-100">
          {users.map((user, index) => (
            <tr key={user.id} className="hover:bg-slate-50 transition-colors">
              <td className="px-5 py-3">{index + 1}</td>

              {user.picture ? (
                <td className="px-5 py-3">
                  <img
                    src={"http://localhost:8080/" + user.picture}
                    alt={user.username}
                    className="w-30 h-30 object-cover"
                  />
                </td>
              ) : (
                <td className="px-5 py-3">No Picture</td>
              )}

              <td className="px-5 py-3 font-medium">{user.username}</td>

              <td className="px-5 py-3">{user.email}</td>

              <td className="px-5 py-3">{user.phone}</td>

              <td className="px-5 py-3">
                <div className="flex justify-end gap-2">
                  <button
                    onClick={() => onEdit(user)}
                    className="px-3 py-1 rounded-lg text-blue-600 hover:bg-blue-50 cursor-pointer"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => onUpload(user)}
                    className="px-3 py-1 rounded-lg text-blue-600 hover:bg-blue-50 cursor-pointer"
                  >
                    Upload Picture
                  </button>

                  <button
                    onClick={() => onDelete(user.id)}
                    className="px-3 py-1 rounded-lg text-red-600 hover:bg-red-50 cursor-pointer"
                  >
                    Hapus
                  </button>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default UserTable;
