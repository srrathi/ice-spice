import xarray as xr
import os

netcdf_file_name = 'FMI-BAL-SEAICE_CONC-L4-NRT-OBS_1674452826620.nc'
csv_file_out = netcdf_file_name[:-3] + '.csv'

ds = xr.open_dataset(netcdf_file_name)
df = ds.to_dataframe()

df.to_csv(csv_file_out)

chunk_size = 40000
def write_chunk(part, lines):
    with open('./tmp/csvs/data_part_'+ str(part) +'.csv', 'w') as f_out:
        f_out.write(header)
        f_out.writelines(lines)
with open("./FMI-BAL-SEAICE_CONC-L4-NRT-OBS_1674452826620.csv", "r") as f:
    count = 0
    header = f.readline()
    lines = []
    for line in f:
        count += 1
        lines.append(line)
        if count % chunk_size == 0:
            write_chunk(count // chunk_size, lines)
            lines = []
    # write remainder
    if len(lines) > 0:
        write_chunk((count // chunk_size) + 1, lines)