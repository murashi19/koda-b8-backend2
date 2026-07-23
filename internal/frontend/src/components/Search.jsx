import { FiSearch } from "react-icons/fi";

function Search({ keyword, onChange, onSubmit }) {
  return (
    <>
      <div className="w-full flex justify-between items-center pl-3 border border-gray-400 rounded-2xl  ">
        <form onSubmit={onSubmit} className="w-full" action="/users">
          <input
            className="w-full outline-none text-md "
            type="text"
            value={keyword}
            onChange={onChange}
            placeholder="Search data by username or email..."
            name="search"
            id=""
          />
        </form>
        <div>
          <button
            type="button"
            className="w-full bg-blue-600 rounded-r-2xl text-white p-3"
          >
            <FiSearch />
          </button>
        </div>
      </div>
    </>
  );
}

export default Search;
