import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../../api/axios";

function Register() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    email: "",
    password: "",
    username: "",
    phone: "",
  });
  const [showPassword, setShowPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    setLoading(true);
    try {
      console.log("Form Data:", formData);
      const response = await api.post("/auth/register", formData);
      console.log("Registration successful:", response.data);
      navigate("/login");
    } catch (error) {
      console.error("Error registering user:", error);
      setError(error.response?.data?.message || "Registrasi gagal. Coba lagi.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-linear-to-br from-slate-50 via-blue-50 to-indigo-100 px-4 py-10">
      <div className="w-full max-w-md">
        {/* Logo / brand */}
        <div className="text-center mb-8">
          <div className="inline-flex items-center justify-center w-14 h-14 rounded-2xl bg-blue-600 text-white text-2xl font-bold mb-4 shadow-lg shadow-blue-200">
            A
          </div>
          <h1 className="text-2xl font-bold text-slate-800">Buat Akun</h1>
          <p className="text-slate-500 text-sm mt-1">
            Isi data di bawah untuk mendaftar
          </p>
        </div>

        <div className="bg-white p-8 rounded-2xl shadow-xl shadow-slate-200/50 border border-slate-100">
          <form onSubmit={handleSubmit} className="space-y-5">
            {error && (
              <div className="text-sm text-red-600 bg-red-50 border border-red-100 rounded-lg px-3 py-2">
                {error}
              </div>
            )}

            <div>
              <label
                htmlFor="username"
                className="block text-sm font-medium text-slate-700 mb-1.5"
              >
                Username
              </label>
              <input
                type="text"
                id="username"
                name="username"
                placeholder="username"
                required
                value={formData.username}
                onChange={handleChange}
                className="w-full px-3.5 py-2.5 border border-slate-200 focus:border-blue-500 focus:ring-2 focus:ring-blue-100 rounded-lg outline-none transition-all text-sm"
              />
            </div>

            <div>
              <label
                htmlFor="email"
                className="block text-sm font-medium text-slate-700 mb-1.5"
              >
                Email
              </label>
              <input
                type="email"
                id="email"
                name="email"
                placeholder="nama@email.com"
                required
                value={formData.email}
                onChange={handleChange}
                className="w-full px-3.5 py-2.5 border border-slate-200 focus:border-blue-500 focus:ring-2 focus:ring-blue-100 rounded-lg outline-none transition-all text-sm"
              />
            </div>

            <div>
              <label
                htmlFor="phone"
                className="block text-sm font-medium text-slate-700 mb-1.5"
              >
                Nomor Telepon
              </label>
              <input
                type="text"
                id="phone"
                name="phone"
                placeholder="08xxxxxxxxxx"
                required
                value={formData.phone}
                onChange={handleChange}
                className="w-full px-3.5 py-2.5 border border-slate-200 focus:border-blue-500 focus:ring-2 focus:ring-blue-100 rounded-lg outline-none transition-all text-sm"
              />
            </div>

            <div>
              <label
                htmlFor="password"
                className="block text-sm font-medium text-slate-700 mb-1.5"
              >
                Password
              </label>
              <div className="relative">
                <input
                  type={showPassword ? "text" : "password"}
                  id="password"
                  name="password"
                  placeholder="••••••••"
                  required
                  value={formData.password}
                  onChange={handleChange}
                  className="w-full px-3.5 py-2.5 border border-slate-200 focus:border-blue-500 focus:ring-2 focus:ring-blue-100 rounded-lg outline-none transition-all text-sm pr-10"
                />
                <button
                  type="button"
                  onClick={() => setShowPassword((prev) => !prev)}
                  className="absolute inset-y-0 right-0 flex items-center pr-3 text-slate-400 hover:text-slate-600 text-xs"
                >
                  {showPassword ? "Sembunyikan" : "Lihat"}
                </button>
              </div>
            </div>

            <button
              type="submit"
              disabled={loading}
              className="w-full bg-blue-600 text-white py-2.5 rounded-lg font-medium hover:bg-blue-700 disabled:opacity-60 disabled:cursor-not-allowed transition-colors cursor-pointer"
            >
              {loading ? "Memproses..." : "Register"}
            </button>
          </form>

          <div className="text-center mt-6 text-sm text-slate-500">
            Sudah punya akun?{" "}
            <a
              href="/login"
              className="text-blue-600 font-medium hover:text-blue-700 no-underline"
            >
              Login
            </a>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Register;
