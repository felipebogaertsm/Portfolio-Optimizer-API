# Author: Felipe Bogaerts de Mattos
# email: felipebogaerts@gmail.com
# License: MIT

"""
Plotter file for the Markowitz optimizer.

Reads from the return_variance_data.json file.
"""

import plotly.graph_objects as go
import pandas as pd

# Forming lists:
dataframe = pd.read_json("return_variance_data.json")
return_array = dataframe["Return"].to_numpy()
variance_array = dataframe["Variance"].to_numpy()

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
