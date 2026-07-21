function Modal({ open, title, children, onClose, width = "max-w-md" }) {
  if (!open) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40">
      <div className={`bg-white rounded-xl shadow-xl w-full ${width} mx-4`}>
        <div className="flex items-center justify-between px-6 py-4">
          <h2 className="text-xl font-semibold">{title}</h2>

          <button
            onClick={onClose}
            className="text-gray-500 hover:text-black text-xl cursor-pointer"
          >
            x
          </button>
        </div>

        <div className="p-6">{children}</div>
      </div>
    </div>
  );
}

export default Modal;
