# Author: Felipe Bogaerts de Mattos
# email: felipebogaerts@gmail.com
# License: MIT

"""
Plotter file for the Markowitz optimizer.

Reads from the return_variance_data.json file.
"""

import plotly.graph_objects as go
import numpy as np
import pandas as pd

# Forming lists:
json_data = pd.read_json("return_variance_data.json")
data_list = json_data.values.tolist()

return_array = np.array([])
variance_array = np.array([])

for row in data_list:
    return_array = np.append(return_array, row[0]["Return"])
    variance_array = np.append(variance_array, row[0]["Variance"])

# Plotting:
fig = go.Figure(
    data=go.Scatter(
        y=return_array * 100,
        x=variance_array,
        mode="markers",
        marker=dict(
            size=10,
            colorscale="Viridis",  # one of plotly colorscales
            showscale=True,
        ),
    )
)

fig.show()
